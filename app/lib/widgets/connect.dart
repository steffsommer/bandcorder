import 'package:bandcorder/services/socket_service.dart';
import 'package:bandcorder/shared/validators.dart';
import 'package:flutter/material.dart';

class Connect extends StatefulWidget {
  const Connect({super.key});

  @override
  State<StatefulWidget> createState() => _ConnectState();
}

class _ConnectState extends State<Connect> {
  final SocketService _socketService = SocketService();

  final _formKey = GlobalKey<FormState>();
  bool _isConnecting = false;
  // String _serverAddress = 'http://10.0.2.2:5000';
  String _serverIP = '10.0.2.2';

  @override
  Widget build(BuildContext context) {
    return Container(
        width: double.infinity,
        height: 220,
        padding: const EdgeInsets.all(30.0),
        decoration: const BoxDecoration(
            borderRadius: BorderRadius.all(Radius.circular(20)),
            color: Colors.green,
            shape: BoxShape.rectangle),
        // child: const Text("Container child"),
        child: Form(
            key: _formKey,
            child: Column(
              children: [
                TextFormField(
                  enabled: !_isConnecting,
                  initialValue: _serverIP,
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                    labelText: 'Server IP Address',
                  ),
                  validator: validateIPv4,
                  onChanged: (value) {
                    setState(() {
                      _serverIP = value;
                    });
                  },
                ),
                const SizedBox(height: 30),
                ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    minimumSize: const Size.fromHeight(50),
                  ),
                  onPressed: mayConnect() ? connectToServer : null,
                  child: _isConnecting
                      ? const CircularProgressIndicator()
                      : const Text('Connect'),
                ),
              ],
            )));
  }

  bool mayConnect() {
    final formValid =
        _formKey.currentState?.validate() ?? validateIPv4(_serverIP) == null;
    return formValid && !_isConnecting;
  }

  connectToServer() {
    setState(() {
      _isConnecting = true;
    });
    _socketService.connect(_serverIP).then((_) {
      setState(() {
        _isConnecting = false;
      });
    }).catchError((_) {
      setState(() {
        _isConnecting = false;
      });
    });
  }
}
