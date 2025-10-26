import 'package:bandcorder/services/connection_cache_service.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/file_service.dart';
import 'package:bandcorder/services/recording_service.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:bandcorder/services/web_socket_service.dart';
import 'package:provider/provider.dart';

import 'package:provider/single_child_widget.dart';

List<SingleChildWidget> get providers {
  return [
    Provider(create: (context) => ToastService()),
    Provider(create: (context) => ConnectionConfig()),
    Provider(create: (context) => ConnectionCacheService()),
    ProxyProvider<ConnectionConfig, FileService>(
        update: (_, connectionConfig, __) => FileService(connectionConfig)),
    ProxyProvider<ConnectionConfig, RecordingService>(
        update: (_, connectionConfig, __) =>
            RecordingService(connectionConfig)),
    ProxyProvider<ConnectionConfig, WebSocketService>(
        update: (_, connectionConfig, __) =>
            WebSocketService(connectionConfig)),
  ];
}
