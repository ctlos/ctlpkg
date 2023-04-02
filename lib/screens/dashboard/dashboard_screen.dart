import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:ctlpkg/responsive.dart';
import 'package:ctlpkg/screens/dashboard/components/package_info_board.dart';
import 'package:flutter/material.dart';

import '../../constants.dart';
import 'components/header.dart';

import 'components/outdated_packages.dart';

class MainContentScreen extends StatelessWidget {
  final StatsResponse stats;
  final DescribeResponse description;
  final bool showOutdated;
  const MainContentScreen({
    Key? key,
    required this.stats,
    required this.description,
    required this.showOutdated,
  }) : super(key: key);

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
                      if (Responsive.isMobile(context) && showOutdated)
                        OutdatedPackages(
                          outdated: stats.outdatedCount.toDouble(),
                          total: stats.packagesCount.toDouble(),
                          outdatedPackagesList: stats.outdatedPackages,
                        ),
                      if (Responsive.isMobile(context) && showOutdated)
                        SizedBox(height: defaultPadding),
                      PackageInfoBoard(
                        description: description,
                      ),
                    ],
                  ),
                ),
                if (!Responsive.isMobile(context))
                  SizedBox(width: defaultPadding),
                if (!Responsive.isMobile(context) && showOutdated)
                  Expanded(
                    flex: 2,
                    child: OutdatedPackages(
                      outdated: stats.outdatedCount.toDouble(),
                      total: stats.packagesCount.toDouble(),
                      outdatedPackagesList: stats.outdatedPackages,
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
