import 'package:bandcorder/style_constants.dart';
import 'package:bandcorder/screens/connect_screen.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Bandcorder',
      theme: ThemeData(
        scaffoldBackgroundColor: StyleConstants.colorSurface1,
        appBarTheme: const AppBarTheme(
          backgroundColor: StyleConstants.colorSurface1,
          foregroundColor: StyleConstants.colorSurface2,
        ),
      ),
      home: const ConnectScreen(),
    );
  }
}
