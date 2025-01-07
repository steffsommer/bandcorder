import 'package:bandcorder/pages/home_page.dart';
import 'package:bandcorder/shared/constants.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    navigatorKey: navigatorKey,
    home: const HomePage(),
    theme: ThemeData(scaffoldBackgroundColor: backgroundColor),
  ));
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Bandcorder',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      // home: const HomePage(),
      home: const HomePage(),
    );
  }
}
