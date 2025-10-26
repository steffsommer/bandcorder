import 'package:bandcorder/routing/routes.dart';
import 'package:bandcorder/screens/record_screen.dart';
import 'package:bandcorder/services/connection_cache_service.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/style_constants.dart';
import 'package:bandcorder/widgets/custom_app_bar.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:bandcorder/widgets/heading.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:go_router/go_router.dart';

import '../services/web_socket_service.dart';
import '../widgets/custom_button.dart';

class ConnectScreen extends StatefulWidget {
  const ConnectScreen({
    super.key,
    required this.socketService,
    required this.connectionCacheService,
    required this.connectionConfig,
  });

  final WebSocketService socketService;
  final ConnectionCacheService connectionCacheService;
  final ConnectionConfig connectionConfig;

  @override
  ConnectScreenState createState() => ConnectScreenState();
}

class ConnectScreenState extends State<ConnectScreen> {
  final _hostController = TextEditingController(text: '10.0.2.2');
  bool _isConnecting = false;
  static bool _isInitialLoad = true;

  @override
  void initState() {
    super.initState();
    _initializeConnection();
  }

  Future<void> _initializeConnection() async {
    await widget.socketService.disconnect();
    final address = await widget.connectionCacheService.queryHost();
    if (address != null) {
      print("Found server address in cache. Connecting right away.");
      _hostController.text = address;
      if (_isInitialLoad) {
        _isInitialLoad = false;
        connect();
      }
    }
  }

  @override
  void dispose() {
    _hostController.dispose();
    super.dispose();
  }

  void connect() async {
    setState(() {
      _isConnecting = true;
    });
    try {
      widget.connectionConfig.host = _hostController.text;
      await widget.socketService.connect();
      widget.connectionCacheService.cacheHost(_hostController.text);
      if (!context.mounted) {
        throw StateError("State is not mounted");
      }
      context.go(Routes.record);
    } finally {
      setState(() {
        _isConnecting = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppBar(),
      body: Padding(
        padding: StyleConstants.padding,
        child: CustomCard(
          child: SingleChildScrollView(
            child: Column(
              children: [
                const Heading(message: "Connect"),
                const SizedBox(height: 60),
                SvgPicture.asset("assets/desktop_pc.svg",
                    semanticsLabel: "Bandcorder logo", height: 160),
                const SizedBox(height: 60),
                const Align(
                  alignment: Alignment.centerLeft,
                  child: Text(
                    "Server IP",
                    style: TextStyle(
                        fontSize: StyleConstants.textSizeNormal,
                        fontWeight: FontWeight.bold),
                  ),
                ),
                TextFormField(
                    controller: _hostController,
                    decoration: const InputDecoration(
                      border: OutlineInputBorder(),
                    )),
                const SizedBox(height: 30),
                SizedBox(
                  height: 60,
                  child: Center(
                      child: _isConnecting
                          ? const CircularProgressIndicator()
                          : CustomButton(
                              color: StyleConstants.colorGreen,
                              onPressed: connect,
                              icon: Icons.start,
                              text: "CONNECT")),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}
