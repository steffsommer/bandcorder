import 'package:bandcorder/widgets/connect.dart';
import 'package:flutter/material.dart';
import '../services/socket_service.dart';
import '../widgets/custom_text_field.dart';
import '../widgets/custom_button.dart';

class HomeV2Page extends StatefulWidget {
  const HomeV2Page({super.key});

  @override
  _HomeV2PageState createState() => _HomeV2PageState();
}

class _HomeV2PageState extends State<HomeV2Page> {
  int _count = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Bandcorder'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            const Connect(),
            // AnimatedSwitcher(
            //     duration: const Duration(milliseconds: 200),
            //     transitionBuilder: (Widget child, Animation<double> animation) {
            //       return ScaleTransition(scale: animation, child: child);
            //     },
            //     child: const Connect()
            //     // child: Text(
            //     //   '$_count',
            //     //   // This key causes the AnimatedSwitcher to interpret this as a "new"
            //     //   // child each time the count changes, so that it will begin its animation
            //     //   // when the count changes.
            //     //   key: ValueKey<int>(_count),
            //     //   style: Theme.of(context).textTheme.headlineMedium,
            //     // ),
            //     ),
            ElevatedButton(
              child: const Text('Increment'),
              onPressed: () {
                setState(() {
                  _count += 1;
                });
              },
            ),
          ],
        ),
      ),
    );
  }
}
