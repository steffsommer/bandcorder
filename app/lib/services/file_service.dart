import 'dart:convert';

import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:http/http.dart' as http;

import '../app_constants.dart';

class FileService {
  final ToastService _toastService;
  final ConnectionConfig _connectionConfig;
  final http.Client _httpClient;

  FileService(
    this._connectionConfig,
    this._toastService,
    this._httpClient,
  );

  Future<void> renameLast(String name) async {
    try {
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/files/renameLast");
      final dto = jsonEncode({"fileName": name});
      final res = await _httpClient
          .post(url, body: dto)
          .timeout(AppConstants.requestTimeout);

      if (res.statusCode != 200) {
        _toastService.toastError("Failed to rename last recording");
        return;
      }
      _toastService.toastSuccess("Recording renamed successfully");
    } catch (e) {
      _toastService.toastError("Failed to rename last recording");
    }
  }
}
