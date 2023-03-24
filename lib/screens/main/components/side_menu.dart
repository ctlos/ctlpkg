import 'package:ctlpkg/constants.dart';
import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:ctlpkg/screens/dashboard/components/header.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

class SideMenu extends StatefulWidget {
  SideMenu({
    Key? key,
  }) : super(key: key);

  @override
  State<SideMenu> createState() => _SideMenuState();
}

class _SideMenuState extends State<SideMenu> {
  final controller = TextEditingController();

  List<Widget> items = [];

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: Column(
        children: [
          DrawerHeader(
            child: Image.asset("assets/images/logo.png"),
          ),
          Padding(
            padding: const EdgeInsets.all(8),
            child: SearchField(
              onSearch: () async {
                var response = await stub.search(SearchRequest(
                  pattern: controller.text,
                ));
                setState(() {
                  for (var element in response.packages) {
                    items.add(DrawerListTile(
                      title: element,
                      svgSrc: "assets/icons/menu_doc.svg",
                      press: () {},
                    ));
                  }
                });
              },
              controller: controller,
            ),
          ),
          ListView(
            children: items,
          ),
        ],
      ),
    );
  }
}

class DrawerListTile extends StatelessWidget {
  const DrawerListTile({
    Key? key,
    // For selecting those three line once press "Command+D"
    required this.title,
    required this.svgSrc,
    required this.press,
  }) : super(key: key);

  final String title, svgSrc;
  final VoidCallback press;

  @override
  Widget build(BuildContext context) {
    return ListTile(
      onTap: press,
      horizontalTitleGap: 0.0,
      leading: SvgPicture.asset(
        svgSrc,
        color: Colors.white54,
        height: 16,
      ),
      title: Text(
        title,
        style: const TextStyle(color: Colors.white54),
      ),
    );
  }
}

class PackageList extends StatelessWidget {
  const PackageList({super.key});

  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
