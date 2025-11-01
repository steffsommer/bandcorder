import 'package:bandcorder/routing/routes.dart';
import 'package:bandcorder/screens/connect_screen.dart';
import 'package:bandcorder/services/connection_cache_service.dart';
import 'package:bandcorder/services/connection_config.dart';
import 'package:bandcorder/services/web_socket_service.dart';
import 'package:bandcorder/widgets/custom_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:go_router/go_router.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';

@GenerateMocks([WebSocketService, ConnectionCacheService, ConnectionConfig])
import 'connect_screen_test.mocks.dart';

void main() {
  late MockWebSocketService mockSocketService;
  late MockConnectionCacheService mockCacheService;
  late MockConnectionConfig mockConfig;

  setUp(() {
    mockSocketService = MockWebSocketService();
    mockCacheService = MockConnectionCacheService();
    mockConfig = MockConnectionConfig();

    when(mockSocketService.disconnect()).thenAnswer((_) async => {});
    when(mockCacheService.queryHost()).thenAnswer((_) async => null);
  });

  Widget createWidget() {
    final router = GoRouter(
      routes: [
        GoRoute(
          path: '/',
          builder: (context, state) => ConnectScreen(
            socketService: mockSocketService,
            connectionCacheService: mockCacheService,
            connectionConfig: mockConfig,
          ),
        ),
        GoRoute(
          path: Routes.record,
          builder: (context, state) => const Scaffold(body: Text('Record')),
        ),
      ],
    );

    return MaterialApp.router(
      routerConfig: router,
    );
  }

  testWidgets('sets connection parameters and connects to entered host',
      (tester) async {
    when(mockSocketService.connect()).thenAnswer((_) async => {});

    await tester.pumpWidget(createWidget());
    await tester.pumpAndSettle();

    await tester.enterText(find.byType(TextFormField), '192.168.1.1');
    await tester.ensureVisible(find.byType(CustomButton));
    await tester.pumpAndSettle();
    await tester.tap(find.byType(CustomButton));
    await tester.pump();

    verify(mockConfig.host = '192.168.1.1').called(1);
    verify(mockSocketService.connect()).called(1);
  });

  testWidgets('shows loading indicator while connecting', (tester) async {
    when(mockSocketService.connect()).thenAnswer((_) async {
      await Future.delayed(const Duration(milliseconds: 100));
    });

    await tester.pumpWidget(createWidget());
    await tester.pumpAndSettle();

    await tester.ensureVisible(find.byType(CustomButton));
    await tester.pumpAndSettle();
    await tester.tap(find.byType(CustomButton));
    await tester.pump();

    expect(find.byType(CircularProgressIndicator), findsOneWidget);
    expect(find.byType(CustomButton), findsNothing);

    await tester.pumpAndSettle();
  });

  testWidgets('caches host on successful connection', (tester) async {
    when(mockSocketService.connect()).thenAnswer((_) async => {});

    await tester.pumpWidget(createWidget());
    await tester.pumpAndSettle();

    await tester.enterText(find.byType(TextFormField), '192.168.1.1');
    await tester.ensureVisible(find.byType(CustomButton));
    await tester.pumpAndSettle();
    await tester.tap(find.byType(CustomButton));
    await tester.pumpAndSettle();

    verify(mockCacheService.cacheHost('192.168.1.1')).called(1);
  });

  testWidgets('auto-connects with cached address on first load',
      (tester) async {
    when(mockCacheService.queryHost()).thenAnswer((_) async => '192.168.1.100');
    when(mockSocketService.connect()).thenAnswer((_) async => {});

    await tester.pumpWidget(createWidget());
    await tester.pumpAndSettle();

    verify(mockConfig.host = '192.168.1.100').called(1);
    verify(mockSocketService.connect()).called(1);
    verify(mockCacheService.cacheHost('192.168.1.100')).called(1);
  });
}
