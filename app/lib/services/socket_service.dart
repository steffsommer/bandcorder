import 'dart:async';

import 'package:bandcorder/pages/recording_control_page.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:bandcorder/shared/constants.dart';
import 'package:flutter/material.dart';
import 'package:socket_io_client/socket_io_client.dart' as sio;
import '../models/connection_status.dart';

/// Connection management and communication with a SocketIO server.
/// TODO: Error on creation of multiple connections, there should only
///       be one active
class SocketService {
  late sio.Socket socket;
  Function(ConnectionStatus)? onConnectionStatusChanged;
  Function(bool)? onRecordingStatusChanged;
  Function(bool)? onErrorCallback;
  final _toastService = ToastService();

  Future<void> connect(String ipv4) {
    var completer = Completer<void>();
    var address = 'http://$ipv4:5000';
    print('Connecting to server');
    socket = sio.io(address, <String, dynamic>{
      'transports': ['websocket', 'polling'],
      'retries': 1,
      'timeout': 5000
    });

    socket.on('connect', (_) {
      _toastService.toastSuccess("Connected to server");
      completer.complete();
      print('connected');
      navigatorKey.currentState?.push(MaterialPageRoute(
        builder: (context) => const RecordingControlPage(),
      ));
      if (onConnectionStatusChanged != null) {
        onConnectionStatusChanged!(ConnectionStatus(
            success: true, message: 'Connection to server established'));
      }
    });

    socket.on('connect_error', (error) {
      _toastService.toastError("Failed to establish connection");
      if (!completer.isCompleted) {
        completer.completeError(error);
      }
      print('Connection Error: $error');
      if (onConnectionStatusChanged != null) {
        onConnectionStatusChanged!(ConnectionStatus(
            success: false, message: 'Connection error: $error'));
      }
      socket.disconnect();
    });

    socket.on('disconnect', (_) {
      _toastService.toastError("Server closed the connection");
      print('disconnected');
      if (onConnectionStatusChanged != null) {
        onConnectionStatusChanged!(
            ConnectionStatus(success: true, message: 'Disconnected'));
      }
    });

    socket.on('RecordingStateChange', (data) {
      print(data);
      print('RecordingStateChanged');
      print(onRecordingStatusChanged != null);
      if (onRecordingStatusChanged != null) {
        print(data['isRecording']);
        try {
          onRecordingStatusChanged!(data['isRecording']);
        } catch (error) {
          print('Key not found');
        }
      }
    });
    return completer.future;
  }

  void disconnect() {
    socket.disconnect();
  }

  void sendStartRecordingEvent() {
    socket.emit('StartRecording');
  }

  void sendStopRecordingEvent() {
    socket.emit('StopRecording');
  }
}
