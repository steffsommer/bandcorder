import 'package:flutter/material.dart';

class Connect extends StatefulWidget {
  const Connect({super.key});

  @override
  State<StatefulWidget> createState() => _ConnectState();
}

class _ConnectState extends State<Connect> {
  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      decoration: const BoxDecoration(
          borderRadius: BorderRadius.all(Radius.circular(2)),
          color: Color.fromRGBO(100, 200, 100, 1),
          shape: BoxShape.rectangle),
      child: const Text("Container child"),
    );
  }
}
