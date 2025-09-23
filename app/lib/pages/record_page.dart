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
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Bandcorder'),
      ),
      body: const Padding(
          padding: EdgeInsets.all(16.0),
          child: CustomCard(
            child: Column(
              children: [
                Timer(isSpinning: true),
                SizedBox(height: 60),
                CustomButton(
                    color: Constants.colorGreen,
                    icon: Icons.play_arrow,
                    text: "START"),
                SizedBox(height: 60)
              ],
            ),
          )),
    );
  }
}
