import 'package:fl_chart/fl_chart.dart';
import 'package:flutter/material.dart';

import '../../../constants.dart';

class Chart extends StatelessWidget {
  final double outdated;
  final double uptodate;
  const Chart({
    Key? key,
    required this.outdated,
    required this.uptodate,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 200,
      child: Stack(
        children: [
          PieChart(
            PieChartData(
              sectionsSpace: 0,
              centerSpaceRadius: 70,
              startDegreeOffset: -90,
              sections: [
                PieChartSectionData(
                  color: primaryColor,
                  value: uptodate,
                  showTitle: false,
                  radius: 25,
                ),
                PieChartSectionData(
                  color: Color(0xFFEE2727),
                  value: outdated,
                  showTitle: false,
                  radius: 22,
                ),
              ],
            ),
          ),
          Positioned.fill(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                SizedBox(height: defaultPadding),
                Text(
                  "$outdated",
                  style: Theme.of(context).textTheme.headlineMedium!.copyWith(
                        color: Colors.white,
                        fontWeight: FontWeight.w600,
                        height: 0.5,
                      ),
                ),
                Text("of $uptodate")
              ],
            ),
          ),
        ],
      ),
    );
  }
}
