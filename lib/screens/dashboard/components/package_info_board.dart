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
                  icon: Icons.receipt,
                  color: Colors.blueGrey,
                ),
                formDataRow(
                  key: "Description",
                  value: description.description,
                  icon: Icons.article,
                  color: Colors.deepOrange,
                ),
                formDataRow(
                  key: "Architecture",
                  value: description.architecture,
                  icon: Icons.computer,
                  color: Colors.orange,
                ),
                formDataRow(
                  value: description.installedSize,
                  key: "Storage size",
                  icon: Icons.sd_card,
                  color: Colors.teal,
                ),
                formDataRow(
                  key: "Official web URL",
                  value: description.url,
                  icon: Icons.web,
                  color: Colors.grey,
                ),
                formDataRow(
                  key: "Licenses",
                  value: description.licenses,
                  icon: Icons.computer,
                  color: Colors.purple,
                ),
                formDataRow(
                  value: description.groups,
                  key: "Groups",
                  icon: Icons.apps,
                  color: Colors.brown,
                ),
                formDataRow(
                  value: description.provides,
                  key: "Provides",
                  icon: Icons.crop_free,
                  color: Colors.indigo,
                ),
                formDataRow(
                  value: description.packager,
                  key: "Packager",
                  icon: Icons.person,
                  color: Colors.blue,
                ),
                formDataRow(
                  value: description.buildDate,
                  key: "Build date",
                  icon: Icons.donut_small,
                  color: Colors.cyan,
                ),
                formDataRow(
                  value: description.conflictsWith,
                  key: "Conflicts",
                  icon: Icons.do_disturb,
                  color: Colors.red,
                ),
                formDataRow(
                  value: description.replaces,
                  key: "Replaces",
                  icon: Icons.crop_3_2,
                  color: Colors.red,
                ),
                formDataRow(
                  value: description.installScript,
                  key: "Has installation script",
                  icon: Icons.code,
                  color: Colors.deepPurple,
                ),
                formDataRow(
                  value: description.validatedBy,
                  key: "Validated by",
                  icon: Icons.check_circle,
                  color: Colors.green,
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
