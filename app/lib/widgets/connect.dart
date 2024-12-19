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
      height: 400,
      padding: EdgeInsets.all(30.0),
      decoration: const BoxDecoration(
          borderRadius: BorderRadius.all(Radius.circular(2)),
          color: Color.fromRGBO(100, 200, 100, 1),
          shape: BoxShape.rectangle),
      // child: const Text("Container child"),
      child: Column(
        children: [
          TextFormField(
            initialValue: 'some default value',
            decoration: const InputDecoration(
              border: OutlineInputBorder(),
              labelText: 'label text',
            ),
          ),
          const SizedBox(height: 30),
          ElevatedButton(
            style: ElevatedButton.styleFrom(
              minimumSize: const Size.fromHeight(
                  50), // fromHeight use double.infinity as width and 40 is the height
            ),
            child: const Text('Connect'),
            onPressed: () {
              // setState(() {
              //   _count += 1;
              // });
            },
          ),
        ],
      ),
    );
  }
}
