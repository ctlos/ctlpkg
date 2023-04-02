import 'package:ctlpkg/controllers/menu_app_controller.dart';
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
  String package = "git";

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      key: context.read<MenuAppController>().scaffoldKey,
      drawer: SideMenu(
        refreshPage: (text) {
          setState(() {
            package = text;
          });
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
                    setState(() {
                      package = text;
                    });
                  },
                ),
              ),
            Expanded(
              // It takes 5/6 part of the screen
              flex: 5,
              child: DashboardScreen(
                key: UniqueKey(),
                package: package,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
