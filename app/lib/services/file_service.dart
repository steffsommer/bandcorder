import 'dart:convert';

import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/toast_service.dart';

import '../app_constants.dart';
import 'package:http/http.dart' as http;

class FileService {
  final _toastService = ToastService();
  final ConnectionConfig _connectionConfig;

  FileService(this._connectionConfig);

  Future<void> renameLast(String name) async {
    try {
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/files/renameLast");
      final dto = jsonEncode({"fileName": name});
      final res =
          await http.post(url, body: dto).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to rename last recording");
      }
      _toastService.toastSuccess("Recording renamed successfully");
    } catch (e) {
      _toastService.toastError("Failed to rename last recording");
    }
  }
}
