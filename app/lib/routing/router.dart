import 'package:bandcorder/globals.dart';
import 'package:bandcorder/routing/routes.dart';
import 'package:bandcorder/screens/connect_screen.dart';
import 'package:bandcorder/screens/metronome_controls_sub_screen.dart';
import 'package:bandcorder/screens/remote_controls_scaffold.dart';
import 'package:bandcorder/widgets/confirmation_dialog.dart';
import 'package:bandcorder/widgets/name_input_dialog.dart';
import 'package:bandcorder/screens/recording_controls_sub_screen.dart';
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
        StatefulShellRoute.indexedStack(
          builder: (context, state, navigationShell) {
            return RemoteControlsScaffold(
              navigationShell: navigationShell,
            );
          },
          branches: [
            StatefulShellBranch(routes: [
              GoRoute(
                  path: Routes.record,
                  builder: (context, state) {
                    return RecordingControlsSubScreen(
                      websocketService: context.read(),
                      recordingService: context.read(),
                      fileService: context.read(),
                      askUserForNewName: showNameInputDialog,
                      askUserForConfirmation: showConfirmationDialog,
                    );
                  }),
            ]),
            StatefulShellBranch(routes: [
              GoRoute(
                  path: Routes.metronome,
                  builder: (context, state) {
                    return MetronomeControlsSubScreen(
                      websocketService: context.read(),
                      metronomeService: context.read(),
                    );
                  })
            ])
          ],
        )
      ],
    );
