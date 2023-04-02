import 'package:ctlpkg/constants.dart';
import 'package:ctlpkg/generated/v1/pacman.pbgrpc.dart';
import 'package:data_table_2/data_table_2.dart';
import 'package:flutter/material.dart';

class PackageInfoBoard extends StatelessWidget {
  final DescribeResponse description;

  const PackageInfoBoard({
    Key? key,
    required this.description,
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
              description.name.toUpperCase(),
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
                formDataRow(
                  key: "Current version",
                  value: description.version,
                  icon: Icons.article,
                  color: Colors.blueGrey,
                ),
                formDataRow(
                  key: "Description",
                  value: description.description,
                  icon: Icons.note,
                  color: Colors.orange,
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

DataRow formDataRow({
  required String key,
  required String value,
  required IconData icon,
  required Color color,
}) {
  return DataRow(
    cells: [
      DataCell(
        Row(
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
            Text(key),
          ],
        ),
      ),
      DataCell(
        Text(value),
      ),
    ],
  );
}
