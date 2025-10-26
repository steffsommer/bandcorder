class ConnectionConfig {
  String host = "";
  int port = 6000;

  String getBaseUrl() {
    return "http://$host:$port";
  }
}
