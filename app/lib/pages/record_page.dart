import 'package:bandcorder/widgets/custom_card.dart';
import 'package:bandcorder/widgets/timer.dart';
import 'package:flutter/material.dart';

class RecordPage extends StatefulWidget {
  const RecordPage({super.key});

  @override
  ConnectPageState createState() => ConnectPageState();
}

class ConnectPageState extends State<RecordPage> {
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
              ],
            ),
          )),
    );
  }
}
