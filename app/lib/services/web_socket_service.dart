import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/routing/routes.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:go_router/go_router.dart';

import '../globals.dart';

const websocketPath = "/ws";
const connectTimeout = Duration(seconds: 3);
const pingFrequency = Duration(milliseconds: 500);

class WebSocketService {
  final _eventsToCallbacks = <Type, List<Function>>{};
  final _toastService = ToastService();
  static const String _intentionalCloseReason = "Disconnect by User";
  WebSocket? _webSocket;
  StreamSubscription<dynamic>? _eventSubscription;
  final ConnectionConfig _connectionConfig;

  WebSocketService(this._connectionConfig);

  /// Establishes a WebSocket connection to the specified host.
  ///
  /// Connects to `ws://[host]:[port][websocketPath]` and starts listening for
  /// incoming events. When events are received, they are deserialized and
  /// dispatched to registered callbacks based on event type.
  ///
  /// The connection will timeout after [connectTimeout] if the server doesn't
  /// respond. Failed events are logged but don't interrupt the connection.
  ///
  /// Throws:
  /// - [TimeoutException] if the server is unreachable within the timeout period
  /// - Other exceptions for connection failures
  Future<void> connect() async {
    try {
      await disconnect();

      final url =
          'ws://${_connectionConfig.host}:${_connectionConfig.port}$websocketPath';
      print('Connecting to $url');

      // Chaining .timeout() behind .connect() is not viable, because the default
      // underlying HTTP client keeps connecting with the default timeout of
      // 30 seconds, which leads to multiple connections being created
      // https://github.com/dart-lang/http/issues/1598
      final httpClient = HttpClient();
      httpClient.connectionTimeout = connectTimeout;
      _webSocket = await WebSocket.connect(url, customClient: httpClient);

      _webSocket!.pingInterval = pingFrequency;

      print('Websocket connection established');

      _eventSubscription = _webSocket!.listen(
        (data) {
          try {
            var event = _convertJsonToEvent(data);
            var callbacks = _eventsToCallbacks[event.runtimeType];
            if (callbacks == null) {
              return;
            }
            for (var cb in callbacks) {
              cb(event);
            }
          } catch (e) {
            // print('Failed to deserialize or process event from JSON $data');
          }
        },
        onDone: () {
          if (_webSocket?.closeCode == WebSocketStatus.normalClosure) {
            return;
          }
          _toastService.toastError("Bandcorder was stopped");
          _onConnectionLoss();
        },
        onError: (error) {
          _toastService.toastError("Unknown connection error");
          _onConnectionLoss();
        },
      );
    } on TimeoutException {
      _toastService.toastError(
          "Server not reachable. Check that Bandcorder is running and the IP is correct");
      rethrow;
    } catch (e) {
      _toastService.toastError("Unknown error while connecting to server");
      rethrow;
    }
  }

  void _onConnectionLoss() {
    navigatorKey.currentContext?.go(Routes.connect);
  }

  /// Registers a callback for events of type [T].
  ///
  /// The [callback] will be invoked whenever an event of the specified type
  /// is received. The type parameter [T] must extend [Event].
  ///
  /// Returns an unsubscribe function that can be called to remove the callback
  /// and stop receiving notifications for this event type.
  void Function() on<T extends Event>(void Function(T) callback) {
    _eventsToCallbacks.putIfAbsent(T, () => <Function>[]);
    _eventsToCallbacks[T]?.add(callback);
    return () => _eventsToCallbacks[T]?.remove(callback);
  }

  Future<void> disconnect() async {
    try {
      await _eventSubscription?.cancel();
      await _webSocket?.close(
          WebSocketStatus.normalClosure, _intentionalCloseReason);
      print('disconnected');
    } on WebSocketException catch (e) {
      print('WebSocket error during close: $e');
    } catch (e) {
      print('Unexpected error: $e');
    }
  }

  static Event _convertJsonToEvent(dynamic json) {
    var jsonStr = json.toString();
    var decodedJson = jsonDecode(jsonStr) as Map<String, dynamic>;
    var idStr = decodedJson["eventId"];
    var eventId = EventId.fromString(idStr);
    var eventData = decodedJson["data"];

    switch (eventId) {
      case EventId.recordingIdle:
        return RecordingIdleEvent();
      case EventId.recordingRunning:
        return RecordingRunningEvent.fromJson(eventData);
    }
  }
}
