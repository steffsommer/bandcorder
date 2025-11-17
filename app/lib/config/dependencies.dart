import 'package:bandcorder/services/connection_cache_service.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/file_service.dart';
import 'package:bandcorder/services/metronome_service.dart';
import 'package:bandcorder/services/recording_service.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:bandcorder/services/web_socket_service.dart';
import 'package:provider/provider.dart';
import 'package:http/http.dart' as http;
import 'package:provider/single_child_widget.dart';

List<SingleChildWidget> get providers {
  return [
    Provider(create: (context) => ToastService()),
    Provider(create: (context) => http.Client()),
    Provider(create: (context) => ConnectionConfig()),
    Provider(create: (context) => ConnectionCacheService()),
    ProxyProvider2<ConnectionConfig, ToastService, FileService>(
      update: (context, connectionConfig, toastService, _) => FileService(
          connectionConfig, toastService, context.read<http.Client>()),
    ),
    ProxyProvider2<ConnectionConfig, ToastService, RecordingService>(
        update: (context, connectionConfig, toastService, __) =>
            RecordingService(
                connectionConfig, toastService, context.read<http.Client>())),
    ProxyProvider2<ConnectionConfig, ToastService, WebSocketService>(
        update: (_, connectionConfig, toastService, __) =>
            WebSocketService(connectionConfig, toastService)),
    ProxyProvider2<ConnectionConfig, ToastService, MetronomeService>(
        update: (context, connectionConfig, toastService, __) =>
            MetronomeService(
                connectionConfig, toastService, context.read<http.Client>())),
  ];
}
