import 'package:flutter/material.dart';

class RecordingButton extends StatelessWidget {
  final String _text;
  final VoidCallback _onPressed;
  final Color _bgColor;
  final Color _textColor;

  const RecordingButton(
      {super.key,
      required String text,
      required void Function() onPressed,
      required Color textColor,
      required Color bgColor,
      required})
      : _onPressed = onPressed,
        _text = text,
        _bgColor = bgColor,
        _textColor = textColor;

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: _onPressed,
      style: ElevatedButton.styleFrom(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5), // <-- Radius
        ),
        minimumSize: const Size(1000, 200),
        backgroundColor: _bgColor,
      ),
      child: Text(
        _text,
        style: TextStyle(fontSize: 40, color: _textColor),
      ),
    );
  }
}
