import 'package:flutter/material.dart';

class RecordingState extends StatefulWidget {
  const RecordingState({super.key});

  @override
  State<StatefulWidget> createState() => _RecordingStateState();
}

class _RecordingStateState extends State<RecordingState> {
  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      height: 400,
      padding: const EdgeInsets.all(30.0),
      decoration: const BoxDecoration(
          borderRadius: BorderRadius.all(Radius.circular(2)),
          color: Color.fromRGBO(200, 100, 100, 1),
          shape: BoxShape.rectangle),
      // child: const Text("Container child"),
      child: const Column(
        children: [
          Text("RECORDING STATE !!!!"),
          SizedBox(height: 30),
        ],
      ),
    );
  }
}
