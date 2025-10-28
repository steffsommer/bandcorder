import 'package:bandcorder/services/recording_service.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';

import 'recording_service_test.mocks.dart';

@GenerateMocks([ConnectionConfig, ToastService, http.Client])
void main() {
  late MockConnectionConfig mockConfig;
  late MockToastService mockToast;
  late MockClient mockClient;
  late RecordingService service;

  setUp(() {
    mockConfig = MockConnectionConfig();
    mockToast = MockToastService();
    mockClient = MockClient();
    service = RecordingService(mockConfig, mockToast, mockClient);

    when(mockConfig.getBaseUrl()).thenReturn('http://test.com');
  });

  group('startRecording', () {
    test('makes POST request to correct URL on success', () async {
      when(mockClient.post(any)).thenAnswer((_) async => http.Response('', 200));

      await service.startRecording();

      verify(mockClient.post(Uri.parse('http://test.com/recording/start'))).called(1);
      verifyNever(mockToast.toastError(any));
    });

    test('shows error toast on non-200 response', () async {
      when(mockClient.post(any)).thenAnswer((_) async => http.Response('', 500));

      await service.startRecording();

      verify(mockToast.toastError('Failed to start recording')).called(1);
    });

    test('shows error toast on exception', () async {
      when(mockClient.post(any)).thenThrow(Exception('Network error'));

      await service.startRecording();

      verify(mockToast.toastError('Failed to start recording')).called(1);
    });
  });

  group('stopRecording', () {
    test('makes POST request to correct URL on success', () async {
      when(mockClient.post(any)).thenAnswer((_) async => http.Response('', 200));

      await service.stopRecording();

      verify(mockClient.post(Uri.parse('http://test.com/recording/stop'))).called(1);
      verifyNever(mockToast.toastError(any));
    });

    test('shows error toast on non-200 response', () async {
      when(mockClient.post(any)).thenAnswer((_) async => http.Response('', 500));

      await service.stopRecording();

      verify(mockToast.toastError('Failed to stop recording')).called(1);
    });

    test('shows error toast on exception', () async {
      when(mockClient.post(any)).thenThrow(Exception('Network error'));

      await service.stopRecording();

      verify(mockToast.toastError('Failed to stop recording')).called(1);
    });
  });

  group('abortRecording', () {
    test('makes POST request to correct URL on success', () async {
      when(mockClient.post(any)).thenAnswer((_) async => http.Response('', 200));

      await service.abortRecording();

      verify(mockClient.post(Uri.parse('http://test.com/recording/abort'))).called(1);
      verifyNever(mockToast.toastError(any));
    });

    test('shows error toast on non-200 response', () async {
      when(mockClient.post(any)).thenAnswer((_) async => http.Response('', 500));

      await service.abortRecording();

      verify(mockToast.toastError('Failed to abort recording')).called(1);
    });

    test('shows error toast on exception', () async {
      when(mockClient.post(any)).thenThrow(Exception('Network error'));

      await service.abortRecording();

      verify(mockToast.toastError('Failed to abort recording')).called(1);
    });
  });
}