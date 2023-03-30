import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:ctlpkg/responsive.dart';
import 'package:ctlpkg/screens/dashboard/components/my_fields.dart';
import 'package:flutter/material.dart';

import '../../constants.dart';
import 'components/header.dart';

import 'components/recent_files.dart';
import 'components/storage_details.dart';

class DashboardScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: SingleChildScrollView(
        primary: false,
        padding: EdgeInsets.all(defaultPadding),
        child: Column(
          children: [
            Header(),
            SizedBox(height: defaultPadding),
            Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Expanded(
                  flex: 5,
                  child: Column(
                    children: [
                      MyFiles(),
                      SizedBox(height: defaultPadding),
                      RecentFiles(),
                      if (Responsive.isMobile(context))
                        SizedBox(height: defaultPadding),
                      if (Responsive.isMobile(context))
                        StarageDetails(
                          outdated: 4,
                          uptodate: 124,
                          outdatedPackagesList: [
                            OutdatedPackage(
                              name: "example",
                              currentVersion: "v1.23.2",
                              latestVersion: "v1.24.1",
                            ),
                            OutdatedPackage(
                              name: "example",
                              currentVersion: "v1.23.2",
                              latestVersion: "v1.24.1",
                            ),
                            OutdatedPackage(
                              name: "example",
                              currentVersion: "v1.23.2",
                              latestVersion: "v1.24.1",
                            ),
                            OutdatedPackage(
                              name: "example",
                              currentVersion: "v1.23.2",
                              latestVersion: "v1.24.1",
                            ),
                          ],
                        ),
                    ],
                  ),
                ),
                if (!Responsive.isMobile(context))
                  SizedBox(width: defaultPadding),
                // On Mobile means if the screen is less than 850 we dont want to show it
                if (!Responsive.isMobile(context))
                  Expanded(
                    flex: 2,
                    child: StarageDetails(
                      outdated: 4,
                      uptodate: 124,
                      outdatedPackagesList: [
                        OutdatedPackage(
                          name: "example",
                          currentVersion: "v1.23.2",
                          latestVersion: "v1.24.1",
                        ),
                        OutdatedPackage(
                          name: "example",
                          currentVersion: "v1.23.2",
                          latestVersion: "v1.24.1",
                        ),
                        OutdatedPackage(
                          name: "example",
                          currentVersion: "v1.23.2",
                          latestVersion: "v1.24.1",
                        ),
                        OutdatedPackage(
                          name: "example",
                          currentVersion: "v1.23.2",
                          latestVersion: "v1.24.1",
                        ),
                      ],
                    ),
                  ),
              ],
            )
          ],
        ),
      ),
    );
  }
}
