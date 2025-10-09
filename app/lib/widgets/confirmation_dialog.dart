import 'package:bandcorder/style_constants.dart';
import 'package:bandcorder/widgets/custom_button.dart';
import 'package:bandcorder/widgets/heading.dart';
import 'package:flutter/material.dart';

Future<bool?> showConfirmationDialog({
  required BuildContext context,
  required String message,
}) {
  return showDialog<bool>(
    context: context,
    builder: (BuildContext context) {
      return AlertDialog(
        title: const Heading(message: "CONFIRM"),
        content: Text(message,
            style: const TextStyle(fontSize: StyleConstants.textSizeBigger)),
        actions: [
          CustomButton(
              color: StyleConstants.colorGreen,
              icon: Icons.play_arrow,
              text: "YES",
              onPressed: () => Navigator.of(context).pop(true)),
          const SizedBox(height: 15),
          CustomButton(
              color: StyleConstants.colorPurple,
              icon: Icons.play_arrow,
              text: "NO",
              onPressed: () => Navigator.of(context).pop(false)),
        ],
      );
    },
  );
}
