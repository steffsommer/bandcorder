import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/services/metronome_service.dart';
import 'package:bandcorder/style_constants.dart';
import 'package:bandcorder/widgets/custom_button.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:flutter/material.dart';
import '../services/web_socket_service.dart';

class MetronomeControlsSubScreen extends StatefulWidget {
  const MetronomeControlsSubScreen({
    super.key,
    required this.websocketService,
    required this.metronomeService,
  });

  final WebSocketService websocketService;
  final MetronomeService metronomeService;

  @override
  MetronomeControlsSubScreenState createState() =>
      MetronomeControlsSubScreenState();
}

class MetronomeControlsSubScreenState
    extends State<MetronomeControlsSubScreen> {
  List<void Function()> cleanupFns = [];
  double _currentBpm = 120.0;
  bool _isMetronomeOn = false;

  @override
  void initState() {
    super.initState();
    cleanupFns = [
      widget.websocketService.on<RecordingRunningEvent>((event) {
        setState(() {
          // TODO: implement metronome logic
        });
      }),
    ];
  }

  @override
  void dispose() {
    super.dispose();
    for (var cleanup in cleanupFns) {
      cleanup();
    }
  }

  void _handleButtonPress() async {
    if (_isMetronomeOn) {
      await widget.metronomeService.stop();
    } else {
      await widget.metronomeService.start();
    }
    setState(() {
      _isMetronomeOn = !_isMetronomeOn;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: StyleConstants.padding,
      child: CustomCard(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              '${_currentBpm.round()} BPM',
              style: const TextStyle(
                fontSize: StyleConstants.textSizeBiggest * 1.5,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: StyleConstants.spacing * 2),
            Slider(
              value: _currentBpm,
              min: 40,
              max: 240,
              onChanged: (value) {
                setState(() {
                  _currentBpm = value;
                });
              },
              onChangeEnd: (value) {
                widget.metronomeService.updateBpm(value.round());
              },
            ),
            const SizedBox(height: StyleConstants.spacing * 2),
            CustomButton(
              color: _isMetronomeOn
                  ? StyleConstants.colorPurple
                  : StyleConstants.colorGreen,
              icon: _isMetronomeOn ? Icons.stop : Icons.play_arrow,
              text: _isMetronomeOn ? 'STOP' : 'START',
              onPressed: _handleButtonPress,
            ),
          ],
        ),
      ),
    );
  }
}
