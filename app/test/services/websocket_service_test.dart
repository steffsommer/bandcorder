import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/services/web_socket_service.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';

import 'websocket_service_test.mocks.dart';

const testTimeout = Duration(milliseconds: 500);

@GenerateMocks([ConnectionConfig, ToastService])
void main() {
  late MockConnectionConfig mockConfig;
  late MockToastService mockToast;
  late WebSocketService service;

  setUp(() {
    mockConfig = MockConnectionConfig();
    mockToast = MockToastService();
    service = WebSocketService(mockConfig, mockToast);
  });

  Future<(HttpServer, Future<WebSocket>)> startMockServer(int port,
      {void Function()? onDone}) async {
    final connectedCompleter = Completer<WebSocket>();
    final httpServer = await HttpServer.bind('127.0.0.1', port);
    WebSocket? serverSocket;
    httpServer.transform(WebSocketTransformer()).listen((ws) {
      serverSocket = ws;
      connectedCompleter.complete(ws);
      ws.listen((data) {}, onDone: onDone);
    });
    addTearDown(() async {
      await service.disconnect(); // Disconnect first
      await serverSocket?.close();
      await httpServer.close(force: true);
    });
    return (httpServer, connectedCompleter.future);
  }

  group('connect', () {
    test('establishes connection successfully', () async {
      final (srv, connected) = await startMockServer(8080);
      when(mockConfig.host).thenReturn('127.0.0.1');
      when(mockConfig.port).thenReturn(8080);

      await service.connect();
      await connected;
    });

    test('shows error on connection failure', () async {
      when(mockConfig.host).thenReturn('192.168.1.255');
      when(mockConfig.port).thenReturn(9999);

      try {
        await service.connect();
        fail('Should have thrown exception');
      } catch (e) {
        expect(e, isA<SocketException>());
        verify(mockToast.toastError(any)).called(1);
      }
    });
  });

  group('disconnect', () {
    test('completes without error when not connected', () async {
      await expectLater(service.disconnect(), completes);
    });

    test('disconnects controlled', () async {
      final disconnectCompleter = Completer<void>();
      final (srv, connected) = await startMockServer(8085,
          onDone: () => disconnectCompleter.complete());
      when(mockConfig.host).thenReturn('127.0.0.1');
      when(mockConfig.port).thenReturn(8085);

      await service.connect();
      await connected;
      await service.disconnect();
      await disconnectCompleter.future.timeout(testTimeout);
    });
  });

  group('event handling', () {
    test('deserializes and emits RecordingIdleEvent', () async {
      final (srv, connected) = await startMockServer(8081);
      when(mockConfig.host).thenReturn('127.0.0.1');
      when(mockConfig.port).thenReturn(8081);

      final completer = Completer<RecordingIdleEvent>();
      service.on<RecordingIdleEvent>((event) => completer.complete(event));

      await service.connect();
      var serverSocket = await connected;

      final eventJson = jsonEncode({'eventId': 'RecordingIdle', 'data': null});
      serverSocket.add(eventJson);

      final event = await completer.future.timeout(testTimeout);
      expect(event, isA<RecordingIdleEvent>());
    });

    test('deserializes and emits RecordingRunningEvent', () async {
      final (srv, connected) = await startMockServer(8082);
      when(mockConfig.host).thenReturn('127.0.0.1');
      when(mockConfig.port).thenReturn(8082);

      final completer = Completer<RecordingRunningEvent>();
      service.on<RecordingRunningEvent>((event) => completer.complete(event));

      await service.connect();
      var serverSocket = await connected;

      final eventJson = jsonEncode({
        'eventId': 'RecordingRunning',
        'data': {'secondsRunning': 123, 'fileName': 'test_file.wav'}
      });
      serverSocket.add(eventJson);

      final event = await completer.future.timeout(testTimeout);
      expect(event, isA<RecordingRunningEvent>());
    });

    test('multiple subscribers receive same event', () async {
      final (srv, connected) = await startMockServer(8083);
      when(mockConfig.host).thenReturn('127.0.0.1');
      when(mockConfig.port).thenReturn(8083);

      final completer = Completer<void>();
      var count1 = 0, count2 = 0;
      service.on<RecordingIdleEvent>((_) {
        count1++;
        if (count1 == 1 && count2 == 1) completer.complete();
      });
      service.on<RecordingIdleEvent>((_) {
        count2++;
        if (count1 == 1 && count2 == 1) completer.complete();
      });

      await service.connect();
      var serverSocket = await connected;

      final eventJson = jsonEncode({'eventId': 'RecordingIdle', 'data': null});
      serverSocket.add(eventJson);

      await completer.future.timeout(testTimeout);
      expect(count1, 1);
      expect(count2, 1);
    });

    test('unsubscribe stops receiving events', () async {
      final (srv, connected) = await startMockServer(8084);
      when(mockConfig.host).thenReturn('127.0.0.1');
      when(mockConfig.port).thenReturn(8084);

      var count = 0;
      final unsubscribe = service.on<RecordingIdleEvent>((_) => count++);

      await service.connect();
      var serverSocket = await connected;

      unsubscribe();

      final eventJson = jsonEncode({'eventId': 'recordingIdle', 'data': null});
      serverSocket.add(eventJson);

      await Future.delayed(testTimeout);
      expect(count, 0);
    });
  });
}
