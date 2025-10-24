import 'dart:convert';

import 'package:bandcorder/services/toast_service.dart';

import '../app_constants.dart';
import 'package:http/http.dart' as http;

class FileService {
  static final FileService instance = FileService();
  String _baseUrl = "";
  final _toastService = ToastService();

  init(String host) {
    _baseUrl = "http://$host:${AppConstants.serverPort}";
  }

  Future<void> renameLast(String name) async {
    _assertBaseUrl();
    try {
      final url = Uri.parse("$_baseUrl/files/renameLast");
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

  _assertBaseUrl() {
    if (_baseUrl == "") {
      throw StateError("RecordingService has not been initialized");
    }
  }
}
