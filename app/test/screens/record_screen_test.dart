import 'dart:async';

import 'package:bandcorder/models/event.dart';
import 'package:bandcorder/screens/record_screen.dart';
import 'package:bandcorder/services/file_service.dart';
import 'package:bandcorder/services/recording_service.dart';
import 'package:bandcorder/services/web_socket_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';

import 'record_screen_test.mocks.dart';

@GenerateMocks([WebSocketService, RecordingService, FileService])
void main() {
  late MockWebSocketService mockWebSocketService;
  late MockRecordingService mockRecordingService;
  late MockFileService mockFileService;
  late void Function(RecordingRunningEvent) onRunningEvent;
  late void Function(RecordingIdleEvent) onIdleEvent;

  setUp(() {
    mockWebSocketService = MockWebSocketService();
    mockRecordingService = MockRecordingService();
    mockFileService = MockFileService();

    when(mockWebSocketService.on(any)).thenAnswer((inv) {
      final callback = inv.positionalArguments[0];
      if (callback is void Function(RecordingRunningEvent)) {
        onRunningEvent = callback;
      } else if (callback is void Function(RecordingIdleEvent)) {
        onIdleEvent = callback;
      }
      return () {};
    });
  });

  Widget createWidget({
    Future<String?> Function(BuildContext)? askUserForNewName,
    Future<bool?> Function(
            {required BuildContext context, required String message})?
        askUserForConfirmation,
  }) {
    return MaterialApp(
      home: RecordScreen(
        websocketService: mockWebSocketService,
        recordingService: mockRecordingService,
        fileService: mockFileService,
        askUserForNewName: askUserForNewName ?? (_) async => null,
        askUserForConfirmation: askUserForConfirmation ??
            ({required context, required message}) async => false,
      ),
    );
  }

  testWidgets('displays idle controls by default', (tester) async {
    await tester.pumpWidget(createWidget());
    expect(find.text('Record'), findsOneWidget);
    expect(find.text('START'), findsOneWidget);
    expect(find.text('RENAME LAST'), findsOneWidget);
  });

  testWidgets('starts recording when START pressed', (tester) async {
    when(mockRecordingService.startRecording()).thenAnswer((_) async => {});

    await tester.pumpWidget(createWidget());

    var startBtn = find.text('START');
    await tester.ensureVisible(startBtn);
    await tester.tap(startBtn);
    await tester.pump();

    verify(mockRecordingService.startRecording()).called(1);
  });

  testWidgets('displays running controls when recording starts',
      (tester) async {
    await tester.pumpWidget(createWidget());
    await tester.pump();

    onRunningEvent(
        RecordingRunningEvent(fileName: 'test.wav', secondsRunning: 10));
    await tester.pump();

    expect(find.text('STOP'), findsOneWidget);
    expect(find.text('ABORT'), findsOneWidget);
    expect(find.text('test.wav'), findsOneWidget);
  });

  testWidgets('stops recording when STOP pressed', (tester) async {
    when(mockRecordingService.stopRecording()).thenAnswer((_) async => {});

    await tester.pumpWidget(createWidget());
    await tester.pump();

    onRunningEvent(
        RecordingRunningEvent(fileName: 'test.wav', secondsRunning: 10));
    await tester.pump();

    var stopBtn = find.text('STOP');
    await tester.ensureVisible(stopBtn);
    await tester.tap(stopBtn);
    await tester.pump();

    verify(mockRecordingService.stopRecording()).called(1);
  });

  testWidgets('returns to idle controls after recording stops', (tester) async {
    await tester.pumpWidget(createWidget());
    await tester.pump();

    onRunningEvent(
        RecordingRunningEvent(fileName: 'test.wav', secondsRunning: 10));
    await tester.pump();
    onIdleEvent(RecordingIdleEvent());
    await tester.pump();

    expect(find.text('START'), findsOneWidget);
    expect(find.text('RENAME LAST'), findsOneWidget);
  });

  testWidgets('shows confirmation dialog before aborting', (tester) async {
    bool confirmationShown = false;
    await tester.pumpWidget(createWidget(
      askUserForConfirmation: ({required context, required message}) async {
        confirmationShown = true;
        return false;
      },
    ));
    await tester.pump();

    onRunningEvent(
        RecordingRunningEvent(fileName: 'test.wav', secondsRunning: 10));
    await tester.pump();

    var abortBtn = find.text('ABORT');
    await tester.ensureVisible(abortBtn);
    await tester.tap(abortBtn);
    await tester.pump();

    expect(confirmationShown, isTrue);
    verifyNever(mockRecordingService.abortRecording());
  });

  testWidgets('aborts recording when confirmed', (tester) async {
    when(mockRecordingService.abortRecording()).thenAnswer((_) async => {});
    await tester.pumpWidget(createWidget(
      askUserForConfirmation: ({required context, required message}) async =>
          true,
    ));
    await tester.pump();

    onRunningEvent(
        RecordingRunningEvent(fileName: 'test.wav', secondsRunning: 10));
    await tester.pump();

    var abortBtn = find.text('ABORT');
    await tester.ensureVisible(abortBtn);
    await tester.tap(abortBtn);
    await tester.pump();

    verify(mockRecordingService.abortRecording()).called(1);
  });

  testWidgets('renames last file when name provided', (tester) async {
    when(mockFileService.renameLast(any)).thenAnswer((_) async => {});

    await tester.pumpWidget(createWidget(
      askUserForNewName: (_) async => 'new_name.wav',
    ));

    var startBtn = find.text('RENAME LAST');
    await tester.ensureVisible(startBtn);
    await tester.tap(startBtn);
    await tester.pump();

    verify(mockFileService.renameLast('new_name.wav')).called(1);
  });

  testWidgets('does not rename when rename dialog closed', (tester) async {
    await tester.pumpWidget(createWidget(
      askUserForNewName: (_) async => null,
    ));

    var renameBtn = find.text('RENAME LAST');
    await tester.ensureVisible(renameBtn);
    await tester.tap(renameBtn);
    await tester.pumpAndSettle();

    verifyNever(mockFileService.renameLast(any));
  });

  testWidgets('shows loading indicator while processing', (tester) async {
    final completer = Completer<void>();
    when(mockRecordingService.startRecording())
        .thenAnswer((_) => completer.future);

    await tester.pumpWidget(createWidget());

    var startBtn = find.text('START');
    await tester.ensureVisible(startBtn);
    await tester.pumpAndSettle();
    await tester.tap(startBtn);
    await tester.pump();

    expect(find.byType(CircularProgressIndicator), findsOneWidget);

    completer.complete();
    await tester.pumpAndSettle();
  });
}
