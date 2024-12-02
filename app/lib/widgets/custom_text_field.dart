import 'package:flutter/material.dart';

class CustomTextField extends StatelessWidget {
  final Function(String) onChanged;
  final String labelText;

  final String? defaultValue; 
  const CustomTextField({super.key, required this.onChanged, required this.labelText, this.defaultValue});

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      initialValue: defaultValue,
      onChanged: onChanged,
      decoration: InputDecoration(
        border: const OutlineInputBorder(),
        labelText: labelText,

      ),
    );
  }
}
