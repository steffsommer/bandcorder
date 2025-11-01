import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'package:bandcorder/widgets/timer.dart';

void main() {
  testWidgets('Timer displays formatted time when secondsRunning is provided',
      (tester) async {
    await tester.pumpWidget(
      const MaterialApp(
        home: Scaffold(
          body: Timer(secondsRunning: 65),
        ),
      ),
    );

    expect(find.textContaining('01:05'), findsOneWidget);
  });

  testWidgets('Timer displays no time text when secondsRunning is null',
      (tester) async {
    await tester.pumpWidget(
      const MaterialApp(
        home: Scaffold(
          body: Timer(),
        ),
      ),
    );

    expect(find.text('0:00'), findsNothing);
    expect(find.textContaining(':'), findsNothing);
  });

  testWidgets('Timer uses custom size', (tester) async {
    await tester.pumpWidget(
      const MaterialApp(
        home: Scaffold(
          body: Timer(size: 100, secondsRunning: 5),
        ),
      ),
    );

    final sizedBox = tester.widget<SizedBox>(find.byType(SizedBox).first);
    expect(sizedBox.width, 100);
    expect(sizedBox.height, 100);
  });
}
