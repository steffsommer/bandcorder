import 'dart:async';

import 'package:bandcorder/app_constants.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:http/http.dart' as http;

class RecordingService {
  static final RecordingService instance = RecordingService();
  final _toastService = ToastService();
  String _baseUrl = "";

  init(String host) {
    _baseUrl = "http://$host:${AppConstants.serverPort}";
  }

  Future<void> startRecording() async {
    _assertBaseUrl();
    try {
      final url = Uri.parse("$_baseUrl/recording/start");
      final res = await http.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to start recording");
      }
    } catch (e) {
      _toastService.toastError("Failed to start recording");
    }
  }

  Future<void> stopRecording() async {
    _assertBaseUrl();
    try {
      final url = Uri.parse("$_baseUrl/recording/stop");
      final res = await http.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to stop recording");
      }
    } catch (e) {
      _toastService.toastError("Failed to stop recording");
    }
  }

  Future<void> abortRecording() async {
    try {
      final url = Uri.parse("$_baseUrl/recording/abort");
      final res = await http.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to abort recording");
      }
    } catch (e) {
      _toastService.toastError("Failed to abort recording");
    }
  }

  _assertBaseUrl() {
    if (_baseUrl == "") {
      throw StateError("RecordingService has not been initialized");
    }
  }
}
