import 'package:flutter/material.dart';

class Constants {
  static const colorSurface1 = Color(0xFF5EB3F5);
  static const colorSurface2 = Color(0xFFFFFFFF);
  static const colorGreen = Color(0xFFa0ff92);
  static const colorYellow = Color(0xFFf8cb46);
  static const colorPurple = Color(0xFFfe90e9);

  static var border = Border.all(color: Colors.black, width: 1);
  static var borderRadius = BorderRadius.circular(25);
  static const padding = EdgeInsets.all(16.0);

  static const textSizeNormal = 17.0;
  static const textSizeBigger = 20.0;
  static const boxShadow = [
    BoxShadow(
      color: Colors.black,
      offset: Offset(2, 2),
      blurRadius: 0,
    ),
  ];
}
