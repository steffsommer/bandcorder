import 'package:flutter/material.dart';
import '../services/socket_service.dart';
import '../widgets/custom_text_field.dart';
import '../widgets/custom_button.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final SocketService _socketService = SocketService();
  String _textFieldValue = 'http://10.0.0.2:5000'; // DefaultWert
  String _connectionStatus = 'disconnected';
  String _recordingStatus = 'stopped';

  @override
  void initState() {
    super.initState();
    _socketService.onConnectionStatusChanged = (status) {
      setState(() {
        _connectionStatus = status;
        _textFieldValue = 'http://10.0.0.2:5000';
      });

    _socketService.onRecordingStatusChanged = (status) {
      setState(() {
        print(status);
        _recordingStatus = status  ? 'started' : 'stopped';
        print(_recordingStatus);
      });
    };

    };
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Bandcorder'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            CustomTextField(
              labelText: 'Bitte IP und Port des Servers eingeben:',
              onChanged: (value) {
                setState(() {
                  _textFieldValue = value;
                });
              },
            ),
            SizedBox(height: 30),
            Text('Connection Status: $_connectionStatus'),
            SizedBox(height: 30),
            if(_connectionStatus == 'disconnected')
            CustomButton(
              text: 'Connect to Server',
              onPressed: () {
                _socketService.connect(_textFieldValue);
              },
            ),
            if(_connectionStatus == 'connected')
            CustomButton(text: 'DisconnectFromServer', 
            onPressed: () {
              _socketService.disconnect();
            }),

            if(_recordingStatus == 'stopped')
            CustomButton(
              text: 'Starte aufnahme',
              onPressed: () {
                _socketService.SendStartRecordingEvent();
              },
            ),

            if(_recordingStatus == 'started')
            CustomButton(
              text: 'Beende Aufnahme',
              onPressed: () {
                _socketService.SendStopRecordingEvent();
              },
            ),
          ],
        ),
      ),
    );
  }
}
