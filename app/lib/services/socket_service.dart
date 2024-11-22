import 'package:socket_io_client/socket_io_client.dart' as IO;

class SocketService {
  late IO.Socket socket;
  Function(String)? onConnectionStatusChanged;
  Function(bool)? onRecordingStatusChanged;

  SocketService({this.onConnectionStatusChanged, this.onRecordingStatusChanged});

  void connect(String url) {


    if(url == '') {
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
        onConnectionStatusChanged!('connected');
      }
    });

    socket.on('connect_error', (error) {
      print('Connection Error: $error');
      if (onConnectionStatusChanged != null) {
        onConnectionStatusChanged!('connection error');
      }
    });

    socket.on('disconnect', (_) {
      print('disconnected');
      if (onConnectionStatusChanged != null) {
        onConnectionStatusChanged!('disconnected');
      }
    });

    socket.on('RecordingStateChange', (data) {
      print(data);
      print('RecordingStateChanged');
      print(onRecordingStatusChanged != null);
      if(onRecordingStatusChanged != null) {
        print(data['recording']);
        onRecordingStatusChanged!(data['recording']);
      }
    });


  }

    void disconnect() {
    socket.disconnect();
  }

  void SendStartRecordingEvent() {
    socket.emit('StartRecording');
  }

  void SendStopRecordingEvent() {
    socket.emit('StopRecording');
  }


}


