import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/widgets/custom_button.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:bandcorder/widgets/timer.dart';
import 'package:flutter/material.dart';

import '../constants.dart';
import '../services/web_socket_service.dart';

class RecordScreen extends StatefulWidget {
  const RecordScreen({super.key});

  @override
  RecordScreenState createState() => RecordScreenState();
}

class RecordScreenState extends State<RecordScreen> {
  static final WebSocketService websocketService = WebSocketService.instance;
  List<void Function()> cleanupFns = [];
  String recordingName = "";
  DateTime? startTime;

  @override
  void initState() {
    super.initState();
    cleanupFns = [
      websocketService.on<RecordingRunningEvent>((event) {
        setState(() {
          recordingName = event.fileName;
          startTime = event.started;
        });
      }),
      websocketService.on<RecordingIdleEvent>((event) {
        setState(() {
          recordingName = "";
          startTime = null;
        });
      })
    ];
  }

  @override
  void dispose() {
    super.dispose();
    for (var cleanup in cleanupFns) {
      cleanup();
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Bandcorder'),
      ),
      body: Padding(
          padding: const EdgeInsets.all(16.0),
          child: CustomCard(
            child: Column(
              children: [
                const SizedBox(height: 30),
                Timer(startTime: startTime),
                const SizedBox(height: 120),
                Column(children: getControls()),
                const SizedBox(height: 30),
              ],
            ),
          )),
    );
  }

  List<Widget> getControls() {
    if (isRunning()) {
      return getRunningControls();
    }
    return getIdleControls();
  }

  List<Widget> getIdleControls() {
    return const [
      CustomButton(
          color: Constants.colorGreen, icon: Icons.play_arrow, text: "START"),
      SizedBox(height: 30),
      CustomButton(
          color: Constants.colorYellow, icon: Icons.edit, text: "RENAME LAST"),
    ];
  }

  List<Widget> getRunningControls() {
    return const [
      CustomButton(
          color: Constants.colorYellow, icon: Icons.pause, text: "STOP"),
      SizedBox(height: 30),
      CustomButton(
          color: Constants.colorPurple, icon: Icons.stop, text: "ABORT"),
    ];
  }

  bool isRunning() {
    return recordingName != "";
  }
}
