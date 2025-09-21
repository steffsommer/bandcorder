import 'dart:async';

import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

const port = 6000;
const websocketPath = "/ws";
const connectTimeout = Duration(seconds: 3);

class WebSocketService {
  final _eventsToCallbacks = <EventId, List<void Function(Event)>>{};
  final toastService = ToastService();

  Future<void> connect(String host) async {
    try {
      final url = Uri.parse('ws://$host:$port$websocketPath');
      print('Connecting to $url');
      final channel = WebSocketChannel.connect(url);
      await channel.ready.timeout(connectTimeout);

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
    } on TimeoutException {
      toastService.toastError(
          "Server not reachable. Check that Bandcorder is running.");
      rethrow;
    } catch (e) {
      toastService.toastError("Unknown error while connecting to server");
      rethrow;
    }
  }

  on(EventId eventId, void Function(Event) callback) {
    _eventsToCallbacks.putIfAbsent(eventId, () => <void Function(Event)>[]);
    _eventsToCallbacks[eventId]?.add(callback);
  }
}
