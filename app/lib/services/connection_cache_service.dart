import 'package:network_info_plus/network_info_plus.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ConnectionCacheService {
  static const _prefix = 'wifi_ip_';

  Future<void> cacheHost(String address) async {
    final wifiName = await _getWifiName();
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('$_prefix$wifiName', address);
  }

  Future<String?> queryHost() async {
    final wifiName = await _getWifiName();
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString('$_prefix$wifiName');
  }

  Future<String> _getWifiName() async {
    if (await Permission.location.request().isDenied) {
      throw Exception("Location permission required");
    }
    final wifiName = await NetworkInfo().getWifiName();
    if (wifiName == null) throw Exception("Not connected to WiFi");
    return wifiName.replaceAll('"', '');
  }
}
