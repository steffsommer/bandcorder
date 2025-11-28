import 'dart:async';

import 'package:bandcorder/app_constants.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:http/http.dart' as http;

import 'connection_config.dart';

class RecordingService {
  final ToastService _toastService;
  final ConnectionConfig _connectionConfig;
  final http.Client _httpClient;

  RecordingService(
      this._connectionConfig, this._toastService, this._httpClient);

  Future<void> startRecording() async {
    try {
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/recording/start");
      final res =
          await _httpClient.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to start recording");
      }
    } catch (e) {
      _toastService.toastError("Failed to start recording");
    }
  }

  Future<void> stopRecording() async {
    try {
      final url = Uri.parse("${_connectionConfig.getBaseUrl()}/recording/stop");
      final res =
          await _httpClient.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Error stopping recording");
      }
    } catch (e) {
      _toastService.toastError("Error stopping recording");
    }
  }

  Future<void> abortRecording() async {
    try {
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/recording/abort");
      final res =
          await _httpClient.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to abort recording");
      }
    } catch (e) {
      _toastService.toastError("Failed to abort recording");
    }
  }
}
