import 'dart:async';

import 'package:bandcorder/app_constants.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:http/http.dart' as http;

import 'connection_config.dart';

class RecordingService {
  final _toastService = ToastService();
  final ConnectionConfig _connectionConfig;

  RecordingService(this._connectionConfig);

  Future<void> startRecording() async {
    try {
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/recording/start");
      final res = await http.post(url).timeout(AppConstants.requestTimeout);
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
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/recording/abort");
      final res = await http.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to abort recording");
      }
    } catch (e) {
      _toastService.toastError("Failed to abort recording");
    }
  }
}
