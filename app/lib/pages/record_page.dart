import 'package:bandcorder/widgets/custom_button.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:bandcorder/widgets/timer.dart';
import 'package:flutter/material.dart';
import '../constants.dart';

class RecordPage extends StatefulWidget {
  const RecordPage({super.key});

  @override
  RecordPageState createState() => RecordPageState();
}

class RecordPageState extends State<RecordPage> {
  final isRunning = true;

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
                const Timer(isSpinning: true),
                const SizedBox(height: 120),
                Column(children: getControls())
              ],
            ),
          )),
    );
  }

  List<Widget> getControls() {
    if (isRunning) {
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
          color: Constants.colorYellow, icon: Icons.edit, text: "START"),
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
}
