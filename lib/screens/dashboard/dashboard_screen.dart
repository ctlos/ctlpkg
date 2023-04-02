import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:ctlpkg/responsive.dart';
import 'package:ctlpkg/screens/dashboard/components/package_info_board.dart';
import 'package:flutter/material.dart';

import '../../constants.dart';
import 'components/header.dart';

import 'components/outdated_packages.dart';

class DashboardScreen extends StatefulWidget {
  final String package;
  const DashboardScreen({
    Key? key,
    required this.package,
  }) : super(key: key);

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  Widget outdatedPackages = Placeholder();
  DescribeResponse description = DescribeResponse();
  bool showOutdated = false;

  setPackageStatistics() async {
    var resp = await stub.stats(StatsRequest());
    setState(() {
      if (resp.outdatedCount > 0) {
        showOutdated = true;
      }
      outdatedPackages = OutdatedPackages(
        outdated: resp.outdatedCount.toDouble(),
        total: resp.packagesCount.toDouble(),
        outdatedPackagesList: resp.outdatedPackages,
      );
    });
  }

  updatePackageInfo(String text) async {
    var resp = await stub.describe(
      DescribeRequest(package: text),
    );
    setState(() {
      description = resp;
    });
  }

  @override
  void initState() {
    super.initState();
    setPackageStatistics();
    updatePackageInfo(widget.package);
  }

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
                        outdatedPackages,
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
                  Expanded(flex: 2, child: outdatedPackages),
              ],
            )
          ],
        ),
      ),
    );
  }
}
