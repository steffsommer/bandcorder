import 'package:bandcorder/widgets/connect.dart';
import 'package:flutter/material.dart';

class HomeV2Page extends StatefulWidget {
  const HomeV2Page({super.key});

  @override
  _HomeV2PageState createState() => _HomeV2PageState();
}

class _HomeV2PageState extends State<HomeV2Page> {
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
          child: Padding(
        padding: EdgeInsets.all(20),
        child: Connect(),
      )),
    );
  }
}
