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
  int _bpm = 120;
  bool _isRunning = false;

  @override
  void initState() {
    super.initState();
    _loadState();
    cleanupFns = [
      widget.websocketService.on<MetronomeStateChangeEvent>((event) {
        setState(() {
          _bpm = event.bpm;
          _isRunning = event.isRunning;
        });
      }),
    ];
  }

  Future<void> _loadState() async {
    final response = await widget.metronomeService.getState();
    setState(() {
      _bpm = response.bpm;
      _isRunning = response.isRunning;
    });
  }

  @override
  void dispose() {
    super.dispose();
    for (var cleanup in cleanupFns) {
      cleanup();
    }
  }

  void _handleButtonPress() async {
    if (_isRunning) {
      await widget.metronomeService.stop();
    } else {
      await widget.metronomeService.start();
    }
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
              '${_bpm.round()} BPM',
              style: const TextStyle(
                fontSize: StyleConstants.textSizeBiggest * 1.5,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: StyleConstants.spacing * 2),
            Slider(
              value: _bpm.toDouble(),
              min: 40,
              max: 240,
              onChanged: (value) {
                setState(() {
                  _bpm = value.toInt();
                });
              },
              onChangeEnd: (value) {
                widget.metronomeService.updateBpm(value.round());
              },
            ),
            const SizedBox(height: StyleConstants.spacing * 2),
            CustomButton(
              color: _isRunning
                  ? StyleConstants.colorPurple
                  : StyleConstants.colorGreen,
              icon: _isRunning ? Icons.stop : Icons.play_arrow,
              text: _isRunning ? 'STOP' : 'START',
              onPressed: _handleButtonPress,
            ),
          ],
        ),
      ),
    );
  }
}
