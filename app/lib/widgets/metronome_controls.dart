import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:flutter/material.dart';

import '../services/web_socket_service.dart';

class MetronomeControls extends StatefulWidget {
  const MetronomeControls({
    super.key,
    required this.websocketService,
  });

  final WebSocketService websocketService;

  @override
  MetronomeControlsState createState() => MetronomeControlsState();
}

class MetronomeControlsState extends State<MetronomeControls> {
  List<void Function()> cleanupFns = [];
  String recordingName = "";
  int? secondsRunning;
  bool _loading = false;

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

  Future<void> _displayLoadingWhile(Future<void> future) async {
    setState(() {
      _loading = true;
    });
    try {
      await future;
    } finally {
      setState(() {
        _loading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return const Padding(
      padding: EdgeInsets.all(16.0),
      child: CustomCard(
        child: Text("Metronome Controls Coming Soon"),
      ),
    );
  }

}
