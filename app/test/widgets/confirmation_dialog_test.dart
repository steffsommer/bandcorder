import 'package:bandcorder/widgets/confirmation_dialog.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('displays message', (tester) async {
    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(
          body: Builder(
            builder: (context) => ElevatedButton(
              onPressed: () => showConfirmationDialog(
                context: context,
                message: 'Delete this item?',
              ),
              child: const Text('Show'),
            ),
          ),
        ),
      ),
    );

    await tester.tap(find.text('Show'));
    await tester.pumpAndSettle();

    expect(find.text('CONFIRM'), findsOneWidget);
    expect(find.text('Delete this item?'), findsOneWidget);
  });

  testWidgets('returns true on YES', (tester) async {
    bool? result;

    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(
          body: Builder(
            builder: (context) => ElevatedButton(
              onPressed: () async {
                result = await showConfirmationDialog(
                  context: context,
                  message: 'Confirm?',
                );
              },
              child: const Text('Show'),
            ),
          ),
        ),
      ),
    );

    await tester.tap(find.text('Show'));
    await tester.pumpAndSettle();

    await tester.tap(find.text('YES'));
    await tester.pumpAndSettle();

    expect(result, isTrue);
  });

  testWidgets('returns false on NO', (tester) async {
    bool? result;

    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(
          body: Builder(
            builder: (context) => ElevatedButton(
              onPressed: () async {
                result = await showConfirmationDialog(
                  context: context,
                  message: 'Confirm?',
                );
              },
              child: const Text('Show'),
            ),
          ),
        ),
      ),
    );

    await tester.tap(find.text('Show'));
    await tester.pumpAndSettle();

    await tester.tap(find.text('NO'));
    await tester.pumpAndSettle();

    expect(result, isFalse);
  });
}
