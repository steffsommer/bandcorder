import 'package:bandcorder/models/recording_status.dart';
import 'package:bandcorder/services/socket_service.dart';
import 'package:bandcorder/widgets/recording_button.dart';
import 'package:bandcorder/widgets/recording_state_info.dart';
import 'package:flutter/material.dart';

class RecordingControlPage extends StatefulWidget {
  const RecordingControlPage({super.key});

  @override
  State<RecordingControlPage> createState() => _RecordingControlPageState();
}

class _RecordingControlPageState extends State<RecordingControlPage> {
  final SocketService _socketService = SocketService();
  RecordingState? _state;

  _RecordingControlPageState() {
    _setupStateChangeUpdates();
  }

  _setupStateChangeUpdates() {
    _socketService.registerStateChangeCallback((state) => {
          setState(() {
            _state = state;
          })
        });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text(
          'Recording Session',
          style: TextStyle(
            color: Colors.white,
          ),
        ),
        backgroundColor: Colors.black,
      ),
      body: Padding(
          padding: const EdgeInsets.all(20),
          child: Column(
            children: [
              RecordingStateInfo(recordingState: _state),
              if (_state == null)
                const CircularProgressIndicator()
              else if (!_state!.isRecording)
                RecordingButton(
                  onPressed: () {
                    _socketService.sendStartRecordingEvent();
                  },
                  text: 'Start recording',
                  bgColor: Colors.green,
                  textColor: Colors.black,
                )
              else
                RecordingButton(
                  onPressed: () {
                    _socketService.sendStopRecordingEvent();
                  },
                  text: 'Stop recording',
                  bgColor: Colors.red,
                  textColor: Colors.white,
                )
            ],
          )),
    );
  }
}
