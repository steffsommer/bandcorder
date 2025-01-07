import 'dart:async';
import 'dart:developer';

import 'package:bandcorder/pages/home_page.dart';
import 'package:bandcorder/pages/recording_control_page.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:bandcorder/shared/constants.dart';
import 'package:flutter/material.dart';
import 'package:socket_io_client/socket_io_client.dart' as sio;
import '../models/recording_status.dart';

/// Connection management and communication with a SocketIO server.
/// TBD: Error on creation of multiple connections, there should only
///       be one active, currently the UI prevents this
class SocketService {
  late sio.Socket socket;
  final List<Function(RecordingState)> _stateChangeCallbacks = [];
  final _toastService = ToastService();
  static final SocketService _singleton = SocketService._internal();
  SocketService._internal();

  factory SocketService() {
    return _singleton;
  }

  registerStateChangeCallback(Function(RecordingState) cb) {
    _stateChangeCallbacks.add(cb);
  }

  Future<void> connect(String ipv4) {
    var completer = Completer<void>();
    var address = 'http://$ipv4:5000';
    log('Connecting to server');
    socket = sio.io(address, <String, dynamic>{
      'transports': ['websocket', 'polling'],
      'timeout': 4000,
    });

    socket.on('connect', (_) {
      _toastService.toastSuccess("Connected to server");

      if (!completer.isCompleted) {
        completer.complete();
      }
      log('connected');
      navigatorKey.currentState?.push(MaterialPageRoute(
        builder: (context) => const RecordingControlPage(),
      ));
    });

    socket.on('connect_error', (error) {
      _toastService.toastError("Failed to establish connection");
      if (!completer.isCompleted) {
        completer.completeError(error);
      }
      log('Connection Error: $error');
      socket.disconnect();
    });

    socket.on('disconnect', (_) {
      _toastService.toastError("Server closed the connection");
      log('disconnected');
      navigatorKey.currentState?.push(MaterialPageRoute(
        builder: (context) => const HomePage(),
      ));
    });

    socket.on('RecordingStateChange', (data) {
      for (var cb in _stateChangeCallbacks) {
        try {
          final state = RecordingState(
              isRecording: data['isRecording'],
              fileName: data['fileName'],
              duration: data['duration']);
          cb(state);
        } catch (error) {
          log("Failed to deserialize RecordingStateDTO, likely wrong key name. Error: $error");
        }
      }
    });
    return completer.future;
  }

  void disconnect() {
    socket.disconnect();
  }

  void sendStartRecordingEvent() {
    log('Sending request to start recording');
    try {
      socket.emit('StartRecording');
    } catch (e) {
      log('Error when starting recording');
    }
  }

  void sendStopRecordingEvent() {
    log('Sending request to stop recording');
    try {
      socket.emit('StopRecording');
    } catch (e) {
      log('Error stopping recording');
    }
  }
}
