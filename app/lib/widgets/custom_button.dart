import 'package:flutter/material.dart';

import '../constants.dart';

class CustomButton extends StatelessWidget {
  final VoidCallback? onPressed;
  final Color color;
  final String text;
  final IconData icon;

  const CustomButton({
    super.key,
    this.onPressed,
    required this.color,
    required this.icon,
    required this.text,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onPressed,
      child: Container(
        width: double.infinity,
        height: 60,
        decoration: BoxDecoration(
          border: Constants.border,
          color: color,
          boxShadow: const [
            BoxShadow(
              color: Colors.black,
              offset: Offset(2, 2),
              blurRadius: 0,
            ),
          ],
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(
              icon,
              size: 42.0,
            ),
            Text(text,
                style: const TextStyle(
                    fontSize: Constants.textSizeBiggest,
                    fontWeight: FontWeight.bold)),
          ],
        ),
      ),
    );
  }
}
