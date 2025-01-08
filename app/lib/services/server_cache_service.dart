import 'dart:async';
import 'dart:convert';
import 'dart:developer';
import 'dart:io';
import 'package:network_info_plus/network_info_plus.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';

/// Persist and retrieve Server IPs for the Wifi the device is currently
/// connected to. Only works for Android at the moment.
class ServerCacheService {
  static const String _fileName = 'storage_data.json';
  final bool useExternalStorage;

  ServerCacheService({this.useExternalStorage = false});

  /// Cache the given Server address for the current WiFi
  Future<void> cacheServerIP(String address) async {
    try {
      final wifiName = await _getWifiName();
      final file = await _getWriteableCacheFile();
      final cache = await _getServerIPCache();
      cache[wifiName] = address;
      final serializedCache = jsonEncode(cache);
      await file.writeAsString(serializedCache);
      log("updated Cache. Added server address $address for WiFi $wifiName");
    } catch (e) {
      log('Error writing to storage: $e');
    }
  }

  Future<String?> queryServerIP() async {
    final wifiName = await _getWifiName();
    final cache = await _getServerIPCache();
    return cache[wifiName];
  }

  Future<String> _getWifiName() async {
    if (!Platform.isAndroid) {
      throw Exception("Failed to retrieve Wifi name. Unsupported platform");
    }
    if (await Permission.location.request().isGranted == false) {
      throw Exception("Failed to retrieve Wifi name. Permission not granted");
    }
    final info = NetworkInfo();
    final wifiName = await info.getWifiName();
    if (wifiName == null) {
      throw Exception("Not connected to Wifi. Cannot cache");
    }
    return wifiName.replaceAll("\"", "");
  }

  Future<Map<String, String>> _getServerIPCache() async {
    try {
      final file = await _getWriteableCacheFile();
      if (!await file.exists()) {
        return {};
      }
      final jsonString = await file.readAsString();
      final Map<String, dynamic> jsonMap = jsonDecode(jsonString);
      return jsonMap.map((key, value) => MapEntry(key, value.toString()));
    } catch (e) {
      log('Error reading from storage: $e');
      return {};
    }
  }

  Future<File> _getWriteableCacheFile() async {
    final path = await _localPath;
    final file = File('$path/$_fileName');
    if (!await file.parent.exists()) {
      await file.parent.create(recursive: true);
    }
    return file;
  }

  Future<String> get _localPath async {
    if (useExternalStorage && Platform.isAndroid) {
      final directory = await getExternalStorageDirectory();
      if (directory == null) {
        // Fall back to application documents directory if external storage is unavailable
        final appDir = await getApplicationDocumentsDirectory();
        return appDir.path;
      }
      return directory.path;
    } else {
      // Use application documents directory
      final directory = await getApplicationDocumentsDirectory();
      return directory.path;
    }
  }
}
