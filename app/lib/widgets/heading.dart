import 'package:flutter/cupertino.dart';

import '../constants.dart';

class Heading extends StatelessWidget {
  final String _message;

  const Heading({super.key, required String message}) : _message = message;

  @override
  Widget build(BuildContext context) {
    return Text(
      _message,
      style: const TextStyle(
        fontSize: Constants.textSizeBiggest,
        fontWeight: FontWeight.bold,
      ),
    );
  }
}
