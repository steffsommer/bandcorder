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

  String _connectionStatus = 'disconnected';
  bool _connectionEstablished = false;
  String _recordingStatus = 'stopped';
  String _ipDefaultValue = 'http://10.0.0.2:5000';
  String _textFieldValue = '';
  @override
  void initState() {
    _textFieldValue = _ipDefaultValue;
    super.initState();
    _socketService.onConnectionStatusChanged = (status) {


      if(status.success == false && status.message != '') {

        final String errorMessage = 'Beim Verbinden zum Server kam es zu einem Fehler: ${status.message}';
        // showDialog<String>(context: context, 
        // builder: (BuildContext context) => AlertDialog(
        //     title: const Text('Fehler beim Verbinden'),
        //     content: Text(errorMessage),
        //     actions: <Widget>[
        //       TextButton(
        //         onPressed: () => Navigator.pop(context, 'OK'),
        //         child: const Text('OK'),
        //       ),
        //     ],
        //   ),
        
        // );

        }


      setState(() {
        print('Set State');
        _connectionStatus = status.message;
        if(status.success) {
          _connectionEstablished = true;
        } else {
          _connectionEstablished = false;
        }

        print(_connectionEstablished);
      });

    _socketService.onRecordingStatusChanged = (status) {

      setState(() {
        _recordingStatus = status  ? 'started' : 'stopped';
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
              defaultValue: _ipDefaultValue,
              labelText: 'Bitte IP und Port des Servers eingeben:',
              onChanged: (value) {
                setState(() {
                  _textFieldValue = value;
                });
              },
            ),
            const SizedBox(height: 30),
            Text('Verbindungsstatus: $_connectionStatus'),
            const SizedBox(height: 30),
            if(_connectionEstablished == false)
            CustomButton(
              text: 'Zum Server verbinden',
              onPressed: () {
                _socketService.connect(_textFieldValue);
              },
            ),
            if(_connectionEstablished == true)
            CustomButton(text: 'Verbindung trennen', 
            onPressed: () {
              _socketService.disconnect();
            }),

            if(_recordingStatus == 'stopped' && _connectionEstablished) 
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
