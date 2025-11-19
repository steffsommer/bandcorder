import 'dart:convert';

import 'package:bandcorder/models/metronome_state.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:http/http.dart' as http;

import '../app_constants.dart';

class MetronomeService {
  final ToastService _toastService;
  final ConnectionConfig _connectionConfig;
  final http.Client _httpClient;

  MetronomeService(
    this._connectionConfig,
    this._toastService,
    this._httpClient,
  );

  Future<void> start() async {
    try {
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/metronome/start");
      final res =
          await _httpClient.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to start metronome");
        return;
      }
    } catch (e) {
      _toastService.toastError("Failed to start metronome");
    }
  }

  Future<void> stop() async {
    try {
      final url = Uri.parse("${_connectionConfig.getBaseUrl()}/metronome/stop");
      final res =
          await _httpClient.post(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to stop metronome");
        return;
      }
    } catch (e) {
      _toastService.toastError("Failed to stop metronome");
    }
  }

  Future<void> updateBpm(int bpm) async {
    try {
      final dto = jsonEncode({"bpm": bpm});
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/metronome/updateBpm");
      final res = await _httpClient
          .post(url, body: dto)
          .timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        _toastService.toastError("Failed to update bpm");
        return;
      }
    } catch (e) {
      _toastService.toastError("Failed to update bpm");
    }
  }

  Future<MetronomeState> getState() async {
    try {
      final url =
          Uri.parse("${_connectionConfig.getBaseUrl()}/metronome/state");
      final res =
          await _httpClient.get(url).timeout(AppConstants.requestTimeout);
      if (res.statusCode != 200) {
        throw ArgumentError("Failed to decode metronome state");
      }
      final json = jsonDecode(res.body);
      return MetronomeState.fromJson(json);
    } catch (e) {
      _toastService.toastError("Failed to query metronome state");
      throw ArgumentError("Failed to decode metronome state");
    }
  }
}
