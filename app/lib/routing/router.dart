import 'package:bandcorder/globals.dart';
import 'package:bandcorder/routing/routes.dart';
import 'package:bandcorder/screens/connect_screen.dart';
import 'package:bandcorder/screens/record_screen.dart';
import 'package:bandcorder/widgets/confirmation_dialog.dart';
import 'package:bandcorder/widgets/rename_last_dialog.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';

GoRouter router() => GoRouter(
      initialLocation: Routes.connect,
      debugLogDiagnostics: true,
      navigatorKey: navigatorKey,
      routes: [
        GoRoute(
          path: Routes.connect,
          builder: (context, state) {
            return ConnectScreen(
              socketService: context.read(),
              connectionCacheService: context.read(),
              connectionConfig: context.read(),
            );
          },
        ),
        GoRoute(
            path: Routes.record,
            builder: (context, state) {
              return RecordScreen(
                websocketService: context.read(),
                recordingService: context.read(),
                fileService: context.read(),
                askUserForNewName: showNameInputDialog,
                askUserForConfirmation: showConfirmationDialog,
              );
            })
      ],
    );
