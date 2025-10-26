import 'package:bandcorder/config/dependencies.dart';
import 'package:bandcorder/routing/router.dart';
import 'package:bandcorder/style_constants.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(
    MultiProvider(
      providers: providers,
      child: const BandcorderApp(),
    ),
  );
}

class BandcorderApp extends StatelessWidget {
  const BandcorderApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      routerConfig: router(),
      title: 'Bandcorder',
      theme: ThemeData(
        scaffoldBackgroundColor: StyleConstants.colorSurface1,
        appBarTheme: const AppBarTheme(
          backgroundColor: StyleConstants.colorSurface1,
          foregroundColor: StyleConstants.colorSurface2,
        ),
      ),
    );
  }
}
