import 'package:flutter/material.dart';

import '../contants.dart';

class CustomButton extends StatelessWidget {
  final List<Widget> children;
  final VoidCallback? onPressed;
  final Color color;

  const CustomButton({
    super.key,
    required this.children,
    this.onPressed,
    required this.color,
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
          children: children,
        ),
      ),
    );
  }
}
