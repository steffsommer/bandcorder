import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/services/file_service.dart';
import 'package:bandcorder/services/recording_service.dart';
import 'package:bandcorder/widgets/custom_button.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:bandcorder/widgets/heading.dart';
import 'package:bandcorder/widgets/timer.dart';
import 'package:flutter/material.dart';

import '../services/web_socket_service.dart';
import '../style_constants.dart';
import '../widgets/custom_app_bar.dart';

class RecordScreen extends StatefulWidget {
  const RecordScreen({
    super.key,
    required this.websocketService,
    required this.recordingService,
    required this.fileService,
    required this.askUserForNewName,
    required this.askUserForConfirmation,
  });

  final WebSocketService websocketService;
  final RecordingService recordingService;
  final FileService fileService;
  final Future<String?> Function(BuildContext) askUserForNewName;
  final Future<bool?> Function(
      {required BuildContext context,
      required String message}) askUserForConfirmation;

  @override
  RecordScreenState createState() => RecordScreenState();
}

class RecordScreenState extends State<RecordScreen> {
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
          recordingName = event.fileName;
          secondsRunning = event.secondsRunning;
        });
      }),
      widget.websocketService.on<RecordingIdleEvent>((event) {
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
    final confirmed = await widget.askUserForConfirmation(
            context: context,
            message: "Are you sure you want to abort the current recording?") ??
        false;
    if (!confirmed) {
      return;
    }
    await _displayLoadingWhile(widget.recordingService.abortRecording());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppBar(),
      body: Padding(
          padding: const EdgeInsets.all(16.0),
          child: CustomCard(
            child: SingleChildScrollView(
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
                        children: getControls(),
                      )),
                  const SizedBox(height: 30),
                ],
              ),
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
          _displayLoadingWhile(widget.recordingService.startRecording());
        },
      ),
      const SizedBox(height: 30),
      CustomButton(
        color: StyleConstants.colorYellow,
        icon: Icons.edit,
        text: "RENAME LAST",
        onPressed: () async {
          var name = await widget.askUserForNewName(context);
          if (name != null) {
            _displayLoadingWhile(widget.fileService.renameLast(name));
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
          _displayLoadingWhile(widget.recordingService.stopRecording());
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
