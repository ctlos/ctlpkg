import 'package:ctlpkg/constants.dart';
import 'package:ctlpkg/controllers/menu_app_controller.dart';
import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:ctlpkg/responsive.dart';
import 'package:ctlpkg/screens/dashboard/dashboard_screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'components/side_menu.dart';

class MainScreen extends StatefulWidget {
  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  DescribeResponse descr = DescribeResponse();
  StatsResponse stats = StatsResponse();

  updateInformation(String text) async {
    var newDescr = await stub.describe(
      DescribeRequest(package: text),
    );
    var newStats = await stub.stats(StatsRequest());
    setState(() {
      descr = newDescr;
      stats = newStats;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      key: context.read<MenuAppController>().scaffoldKey,
      drawer: SideMenu(
        refreshPage: (text) {
          updateInformation(text);
        },
      ),
      body: SafeArea(
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // We want this side menu only for large screen
            if (Responsive.isDesktop(context))
              Expanded(
                // default flex = 1
                // and it takes 1/6 part of the screen
                child: SideMenu(
                  refreshPage: (text) {
                    updateInformation(text);
                  },
                ),
              ),
            Expanded(
              // It takes 5/6 part of the screen
              flex: 5,
              child: AnimatedSwitcher(
                duration: Duration(milliseconds: 644),
                child: MainContentScreen(
                  key: UniqueKey(),
                  description: descr,
                  stats: stats,
                  showOutdated: stats.outdatedCount > 0,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
