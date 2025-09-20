import 'package:bandcorder/models/event.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

const port = 6000;
const websocketPath = "/ws";

class WebSocketService {
  final _eventsToCallbacks = <EventId, List<void Function(Event)>>{};

  Future<void> connect(String host) async {
    final url = Uri.parse('ws://$host:$port$websocketPath');
    print('Connecting to $url');
    final channel = WebSocketChannel.connect(url);
    await channel.ready;
    print('Websocket connection established');

    channel.stream.listen((message) {
      var event = Event.fromJson(message);
      var callbacks = _eventsToCallbacks[event.id];
      if (callbacks == null) {
        return;
      }
      for (var cb in callbacks) {
        cb(event);
      }
    });
  }

  on(EventId eventId, void Function(Event) callback) {
    _eventsToCallbacks.putIfAbsent(eventId, () => <void Function(Event)>[]);
    _eventsToCallbacks[eventId]?.add(callback);
  }
}
//   socket = IO.io(url, <String, dynamic>{
//     'transports': ['websocket', 'polling'],
//   });
//
//   socket.on('connect', (_) {
//     print('connected');
//     if (onConnectionStatusChanged != null) {
//       onConnectionStatusChanged!(ConnectionStatus(success: true, message: 'Erfolgreich verbunden!'));
//     }
//   });
//
//   socket.on('connect_error', (error) {
//     print('Connection Error: $error');
//     if (onConnectionStatusChanged != null) {
//       onConnectionStatusChanged!(ConnectionStatus(success: false, message: 'Fehler beim Verbinden: $error'));
//     }
//     socket.close();
//
//   });
//
//   socket.on('disconnect', (_) {
//     print('disconnected');
//     if (onConnectionStatusChanged != null) {
//       onConnectionStatusChanged!(ConnectionStatus(success: true, message: 'Erfolgreich getrennt!'));
//     }
//   });
//
//   socket.on('RecordingStateChange', (data) {
//     print(data);
//     print('RecordingStateChanged');
//     print(onRecordingStatusChanged != null);
//     if(onRecordingStatusChanged != null) {
//       print(data['isRecording']);
//       try {
//         onRecordingStatusChanged!(data['Recording']);
//         // onRecordingStatusChanged!(data['isRecording']);
//       } catch(error) {
//         print('Key not found');
//       }
//
//     }
//   });
//
//
// }
//
//   void disconnect() {
//   socket.disconnect();
// }

// void SendStartRecordingEvent() {
//   socket.emit('StartRecording');
// }

// void SendStopRecordingEvent() {
//   socket.emit('StopRecording');
// }
