import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../style_constants.dart';
import '../widgets/custom_app_bar.dart';

class RemoteControlsScaffold extends StatelessWidget {
  const RemoteControlsScaffold({
    super.key,
    required this.navigationShell,
  });

  final StatefulNavigationShell navigationShell;

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
        length: 2,
        child: Scaffold(
          appBar: CustomAppBar(
              bottom: TabBar(
                  onTap: (index) => navigationShell.goBranch(index),
                  labelStyle: const TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                  labelColor: Colors.black,
                  unselectedLabelColor: Colors.grey,
                  indicator: const BoxDecoration(
                    border: Border(
                      bottom: BorderSide(
                        color: StyleConstants.colorGreen,
                        width: 4,
                      ),
                    ),
                  ),
                  indicatorSize: TabBarIndicatorSize.tab,
                  tabs: const [
                Tab(
                  child: Row(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      Icon(Icons.mic),
                      SizedBox(width: 8),
                      Text('RECORD'),
                    ],
                  ),
                ),
                Tab(
                  child: Row(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      Icon(Icons.electric_meter_rounded),
                      SizedBox(width: 8),
                      Text('METRONOME'),
                    ],
                  ),
                ),
              ])),
          body: navigationShell,
        ));
  }
}
