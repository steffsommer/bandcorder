import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:bandcorder/app_constants.dart';
import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:flutter/material.dart';

import '../globals.dart';
import '../screens/connect_screen.dart';

const websocketPath = "/ws";
const connectTimeout = Duration(seconds: 3);
const pingFrequency = Duration(milliseconds: 500);

class WebSocketService {
  final _eventsToCallbacks = <Type, List<Function>>{};
  final _toastService = ToastService();
  static final WebSocketService instance = WebSocketService._();
  WebSocket? _webSocket;

  WebSocketService._();

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
  Future<void> connect(String host) async {
    try {
      final url = 'ws://$host:${AppConstants.serverPort}$websocketPath';
      print('Connecting to $url');

      _webSocket = await WebSocket.connect(url).timeout(connectTimeout);
      _webSocket!.pingInterval = pingFrequency;

      print('Websocket connection established');

      _webSocket!.listen(
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
            print('Failed to deserialize or process event from JSON $data');
          }
        },
        onDone: () {
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
    navigatorKey.currentState?.pushAndRemoveUntil(
      MaterialPageRoute(builder: (context) => const ConnectScreen()),
      (route) => false,
    );
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
