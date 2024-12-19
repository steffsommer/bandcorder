import 'package:socket_io_client/socket_io_client.dart' as IO;
import '../models/connection_status.dart';

class SocketService {
  late IO.Socket socket;
  Function(ConnectionStatus)? onConnectionStatusChanged;
  Function(bool)? onRecordingStatusChanged;
  Function(bool)? onErrorCallback;

  SocketService(
      {this.onConnectionStatusChanged, this.onRecordingStatusChanged});

  void connect(String url) {
    if (url == '') {
      print('Bitte zuerst eine Url eingeben!');
      return;
    }
    print('Start Connect');

    socket = IO.io(url, <String, dynamic>{
      'transports': ['websocket', 'polling'],
    });

    socket.on('connect', (_) {
      print('connected');
      if (onConnectionStatusChanged != null) {
        onConnectionStatusChanged!(
            ConnectionStatus(success: true, message: 'Erfolgreich verbunden!'));
      }
    });

    socket.on('connect_error', (error) {
      print('Connection Error: $error');
      if (onConnectionStatusChanged != null) {
        onConnectionStatusChanged!(ConnectionStatus(
            success: false, message: 'Fehler beim Verbinden: $error'));
      }
      socket.close();
    });

    socket.on('disconnect', (_) {
      print('disconnected');
      if (onConnectionStatusChanged != null) {
        onConnectionStatusChanged!(
            ConnectionStatus(success: true, message: 'Erfolgreich getrennt!'));
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
          // onRecordingStatusChanged!(data['isRecording']);
        } catch (error) {
          print('Key not found');
        }
      }
    });
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
