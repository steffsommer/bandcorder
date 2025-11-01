import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'package:bandcorder/widgets/custom_button.dart';

void main() {
  testWidgets('renders with text and icon', (tester) async {
    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(
          body: CustomButton(
            color: Colors.blue,
            icon: Icons.add,
            text: 'Test Button',
            onPressed: () {},
          ),
        ),
      ),
    );

    expect(find.text('Test Button'), findsOneWidget);
    expect(find.byIcon(Icons.add), findsOneWidget);
  });

  testWidgets('calls onPressed when tapped', (tester) async {
    var pressed = false;

    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(
          body: CustomButton(
            color: Colors.blue,
            icon: Icons.add,
            text: 'Test',
            onPressed: () => pressed = true,
          ),
        ),
      ),
    );

    await tester.tap(find.byType(CustomButton));
    expect(pressed, isTrue);
  });

  testWidgets('handles null onPressed', (tester) async {
    await tester.pumpWidget(
      const MaterialApp(
        home: Scaffold(
          body: CustomButton(
            color: Colors.blue,
            icon: Icons.add,
            text: 'Test',
          ),
        ),
      ),
    );

    await tester.tap(find.byType(CustomButton));
    // Should not throw
  });
}
