import 'package:bandcorder/contants.dart';
import 'package:flutter/material.dart';
import 'pages/connect_page.dart';

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
        scaffoldBackgroundColor: Constants.colorSurface1,
        appBarTheme: const AppBarTheme(
          backgroundColor: Constants.colorSurface1,
          foregroundColor: Constants.colorSurface2,
        ),
      ),
      home: const ConnectPage(),
    );
  }
}
