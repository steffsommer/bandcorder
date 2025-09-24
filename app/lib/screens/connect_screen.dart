import 'package:bandcorder/style_constants.dart';
import 'package:bandcorder/screens/record_screen.dart';
import 'package:bandcorder/widgets/custom_app_bar.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:bandcorder/widgets/heading.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

import '../services/web_socket_service.dart';
import '../widgets/custom_button.dart';

class ConnectScreen extends StatefulWidget {
  const ConnectScreen({super.key});

  @override
  ConnectScreenState createState() => ConnectScreenState();
}

class ConnectScreenState extends State<ConnectScreen> {
  final WebSocketService _socketService = WebSocketService.instance;
  String _textFieldValue = '10.0.2.2';
  bool _isConnecting = false;

  void connect() async {
    setState(() {
      _isConnecting = true;
    });

    try {
      await _socketService.connect(_textFieldValue);
      if (!context.mounted) {
        throw StateError("State is not mounted");
      }
      Navigator.push(
        context,
        MaterialPageRoute<void>(
          builder: (context) => const RecordScreen(),
        ),
      );
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
                  initialValue: _textFieldValue,
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                  ),
                  onChanged: (value) {
                    setState(() {
                      _textFieldValue = value;
                    });
                  }),
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
    );
  }
}
