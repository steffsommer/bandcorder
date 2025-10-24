import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/services/file_service.dart';
import 'package:bandcorder/services/recording_service.dart';
import 'package:bandcorder/widgets/confirmation_dialog.dart';
import 'package:bandcorder/widgets/custom_button.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:bandcorder/widgets/heading.dart';
import 'package:bandcorder/widgets/rename_last_dialog.dart';
import 'package:bandcorder/widgets/timer.dart';
import 'package:flutter/material.dart';

import '../style_constants.dart';
import '../services/web_socket_service.dart';
import '../widgets/custom_app_bar.dart';

class RecordScreen extends StatefulWidget {
  const RecordScreen({super.key});

  @override
  RecordScreenState createState() => RecordScreenState();
}

class RecordScreenState extends State<RecordScreen> {
  static final WebSocketService websocketService = WebSocketService.instance;
  static final RecordingService recordingService = RecordingService.instance;
  static final FileService fileService = FileService.instance;

  List<void Function()> cleanupFns = [];
  String recordingName = "";
  int? secondsRunning;
  bool _loading = false;

  @override
  void initState() {
    super.initState();
    cleanupFns = [
      websocketService.on<RecordingRunningEvent>((event) {
        setState(() {
          recordingName = event.fileName;
          secondsRunning = event.secondsRunning;
        });
      }),
      websocketService.on<RecordingIdleEvent>((event) {
        setState(() {
          recordingName = "";
          secondsRunning = null;
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

  Future<void> _confirmAbortLoading(BuildContext context) async {
    final confirmed = await showConfirmationDialog(
            context: context,
            message: "Are you sure you want to abort the current recording?") ??
        false;
    if (!confirmed) {
      return;
    }
    await _displayLoadingWhile(recordingService.abortRecording());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppBar(),
      body: Padding(
          padding: const EdgeInsets.all(16.0),
          child: CustomCard(
            child: Column(
              children: [
                const Heading(message: "Record"),
                const SizedBox(height: 30),
                Timer(secondsRunning: secondsRunning),
                SizedBox(
                  height: 120,
                  child: recordingName == ""
                      ? null
                      : Center(
                          child: Row(
                            children: [
                              const Icon(
                                Icons.audio_file,
                                size: 30.0,
                              ),
                              Text(recordingName,
                                  style: const TextStyle(
                                      fontSize: StyleConstants.textSizeBigger,
                                      fontWeight: FontWeight.bold)),
                            ],
                          ),
                        ),
                ),
                AnimatedSwitcher(
                    duration: const Duration(milliseconds: 200),
                    child: Column(
                      key: ValueKey('${_loading}_${isRunning()}'),
                      children: getControls(), // Combined key
                    )),
                // _loading
                //     ? const CircularProgressIndicator()
                //     : Column(children: getControls()),
                const SizedBox(height: 30),
              ],
            ),
          )),
    );
  }

  List<Widget> getControls() {
    if (_loading) {
      return [const CircularProgressIndicator()];
    } else if (isRunning()) {
      return getRunningControls();
    }
    return getIdleControls();
  }

  List<Widget> getIdleControls() {
    return [
      CustomButton(
        color: StyleConstants.colorGreen,
        icon: Icons.play_arrow,
        text: "START",
        onPressed: () {
          _displayLoadingWhile(recordingService.startRecording());
        },
      ),
      const SizedBox(height: 30),
      CustomButton(
        color: StyleConstants.colorYellow,
        icon: Icons.edit,
        text: "RENAME LAST",
        onPressed: () async {
          var name = await showNameInputDialog(context);
          if (name != null) {
            _displayLoadingWhile(fileService.renameLast(name));
          }
        },
      ),
    ];
  }

  List<Widget> getRunningControls() {
    return [
      CustomButton(
        color: StyleConstants.colorYellow,
        icon: Icons.pause,
        text: "STOP",
        onPressed: () {
          _displayLoadingWhile(recordingService.stopRecording());
        },
      ),
      const SizedBox(height: 30),
      CustomButton(
        color: StyleConstants.colorPurple,
        icon: Icons.stop,
        text: "ABORT",
        onPressed: () {
          _confirmAbortLoading(context);
        },
      ),
    ];
  }

  bool isRunning() {
    return recordingName != "";
  }
}
