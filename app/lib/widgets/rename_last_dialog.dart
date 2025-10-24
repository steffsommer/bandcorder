import 'package:bandcorder/style_constants.dart';
import 'package:bandcorder/widgets/custom_button.dart';
import 'package:bandcorder/widgets/heading.dart';
import 'package:flutter/material.dart';

Future<String?> showNameInputDialog(BuildContext context) {
  final controller = TextEditingController();
  final nameRegex = RegExp(r'^[a-zA-Z0-9._-]+$');

  return showDialog<String>(
    context: context,
    builder: (BuildContext context) {
      return StatefulBuilder(
        builder: (context, setState) {
          final name = controller.text;
          final isValid = name.isNotEmpty && nameRegex.hasMatch(name);
          final showError = name.isNotEmpty && !nameRegex.hasMatch(name);

          return AlertDialog(
            title: const Heading(message: "ENTER NAME"),
            content: TextField(
              controller: controller,
              onChanged: (_) => setState(() {}),
              decoration: InputDecoration(
                hintText: "File name",
                errorText: showError
                    ? "The file name contains characters that are not allowed"
                    : null,
              ),
              style: const TextStyle(fontSize: StyleConstants.textSizeBigger),
            ),
            actions: [
              CustomButton(
                color: isValid ? StyleConstants.colorGreen : Colors.grey,
                icon: Icons.save,
                text: "SAVE",
                onPressed:
                    isValid ? () => Navigator.of(context).pop(name) : null,
              ),
              const SizedBox(height: 15),
              CustomButton(
                color: StyleConstants.colorPurple,
                icon: Icons.close,
                text: "CANCEL",
                onPressed: () => Navigator.of(context).pop(null),
              ),
            ],
          );
        },
      );
    },
  );
}
