import 'dart:convert';
import 'package:bandcorder/services/file_service.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/toast_service.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'file_service_test.mocks.dart';

@GenerateMocks([ConnectionConfig, ToastService, http.Client])
void main() {
  late MockConnectionConfig mockConfig;
  late MockToastService mockToast;
  late MockClient mockClient;
  late FileService service;

  setUp(() {
    mockConfig = MockConnectionConfig();
    mockToast = MockToastService();
    mockClient = MockClient();
    service = FileService(mockConfig, mockToast, mockClient);

    when(mockConfig.getBaseUrl()).thenReturn('http://test.com');
  });

  group('renameLast', () {
    test('shows success toast on 200 response', () async {
      when(mockClient.post(any, body: anyNamed('body')))
          .thenAnswer((_) async => http.Response('', 200));

      await service.renameLast('test.mp3');

      verify(mockClient.post(
        Uri.parse('http://test.com/files/renameLast'),
        body: jsonEncode({'fileName': 'test.mp3'}),
      )).called(1);
      verify(mockToast.toastSuccess('Recording renamed successfully')).called(1);
      verifyNever(mockToast.toastError(any));
    });

    test('shows error toast on non-200 response', () async {
      when(mockClient.post(any, body: anyNamed('body')))
          .thenAnswer((_) async => http.Response('error', 500));

      await service.renameLast('test.mp3');

      verify(mockClient.post(
        Uri.parse('http://test.com/files/renameLast'),
        body: jsonEncode({'fileName': 'test.mp3'}),
      )).called(1);
      verify(mockToast.toastError('Failed to rename last recording')).called(1);
      verifyNever(mockToast.toastSuccess(any));
    });

    test('shows error toast on exception', () async {
      when(mockClient.post(any, body: anyNamed('body')))
          .thenThrow(Exception('Network error'));

      await service.renameLast('test.mp3');

      verify(mockToast.toastError('Failed to rename last recording')).called(1);
    });
  });
}
