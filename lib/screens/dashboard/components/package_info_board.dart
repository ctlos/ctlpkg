import 'package:ctlpkg/constants.dart';
import 'package:data_table_2/data_table_2.dart';
import 'package:flutter/material.dart';

class PackageInfoBoard extends StatelessWidget {
  final String name;
  const PackageInfoBoard({
    Key? key,
    required this.name,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(defaultPadding),
      decoration: BoxDecoration(
        color: secondaryColor,
        borderRadius: const BorderRadius.all(Radius.circular(10)),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Center(
            child: Text(
              name.toUpperCase(),
              style: Theme.of(context).textTheme.titleMedium,
            ),
          ),
          SizedBox(
            width: double.infinity,
            child: DataTable2(
              columnSpacing: defaultPadding,
              minWidth: 600,
              columns: [
                DataColumn(
                  label: Text("Property"),
                ),
                DataColumn(
                  label: Text("Value"),
                ),
              ],
              rows: [
                DataRow(cells: [
                  DataCell(
                    PropertyNameWidget(
                      icon: Icons.assignment_outlined,
                      name: "Version",
                      color: Colors.yellow,
                    ),
                  ),
                  DataCell(Text("Vupsen")),
                ]),
                DataRow(cells: [
                  DataCell(
                    PropertyNameWidget(
                      icon: Icons.usb,
                      name: "Stuff",
                      color: Colors.red,
                    ),
                  ),
                  DataCell(Text("Pupsen")),
                ]),
                DataRow(cells: [
                  DataCell(
                    PropertyNameWidget(
                      icon: Icons.airline_seat_legroom_reduced,
                      name: "Keks",
                      color: Colors.orange,
                    ),
                  ),
                  DataCell(Text("Mopsel")),
                ]),
                DataRow(cells: [
                  DataCell(
                    PropertyNameWidget(
                      icon: Icons.build_circle_rounded,
                      name: "Ikrop",
                      color: Colors.purple,
                    ),
                  ),
                  DataCell(Text("Keka")),
                ]),
                DataRow(cells: [
                  DataCell(
                    PropertyNameWidget(
                      icon: Icons.call_missed_rounded,
                      name: "Prekl",
                      color: Colors.blue,
                    ),
                  ),
                  DataCell(Text("asdass")),
                ]),
                DataRow(cells: [
                  DataCell(
                    PropertyNameWidget(
                      icon: Icons.brightness_1_rounded,
                      name: "Nanii",
                      color: Colors.blue,
                    ),
                  ),
                  DataCell(Text("12312aa")),
                ]),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

class PropertyNameWidget extends StatelessWidget {
  final IconData icon;
  final String name;
  final Color color;
  const PropertyNameWidget({
    Key? key,
    required this.icon,
    required this.name,
    required this.color,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      children: [
        Container(
          decoration: BoxDecoration(
            color: color,
            borderRadius: const BorderRadius.all(Radius.circular(4)),
          ),
          child: Padding(
            padding: const EdgeInsets.all(2.0),
            child: Icon(
              icon,
              size: 18,
              color: Colors.white,
            ),
          ),
        ),
        Padding(padding: EdgeInsets.all(defaultPadding)),
        Text(name),
      ],
    );
  }
}
