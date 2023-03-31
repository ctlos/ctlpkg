import 'package:ctlpkg/controllers/menu_app_controller.dart';
import 'package:ctlpkg/responsive.dart';
import 'package:ctlpkg/screens/dashboard/components/add_package.dart';
import 'package:ctlpkg/screens/dashboard/components/login_screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../constants.dart';

class Header extends StatelessWidget {
  const Header({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        if (!Responsive.isDesktop(context))
          IconButton(
            icon: Icon(Icons.menu),
            onPressed: context.read<MenuAppController>().controlMenu,
          ),
        if (!Responsive.isMobile(context))
          Text(
            "CtlOS package repository",
            style: Theme.of(context).textTheme.titleLarge,
          ),
        if (!Responsive.isMobile(context))
          Spacer(flex: Responsive.isDesktop(context) ? 2 : 1),
        ProfileCard(),
      ],
    );
  }
}

class ProfileCard extends StatefulWidget {
  const ProfileCard({
    Key? key,
  }) : super(key: key);

  @override
  State<ProfileCard> createState() => _ProfileCardState();
}

class _ProfileCardState extends State<ProfileCard> {
  trySetAuthorized() async {
    var prefs = await SharedPreferences.getInstance();
    var token = prefs.getString("token") ?? "";
    if (token != "") {
      setState(() {
        placeholder = AuthorizedActions(
          add: () {
            showAddPackage(context);
          },
          update: () {
            print("yo");
          },
        );
      });
    }
  }

  Widget placeholder = UnauthorizedWidget(loggedCallback: () {});

  @override
  void initState() {
    placeholder = UnauthorizedWidget(loggedCallback: () async {
      await trySetAuthorized();
    });
    super.initState();
    trySetAuthorized();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.only(left: defaultPadding),
      padding: EdgeInsets.symmetric(
        horizontal: defaultPadding,
        vertical: defaultPadding / 2,
      ),
      child: placeholder,
    );
  }
}

class UnauthorizedWidget extends StatelessWidget {
  final void Function() loggedCallback;
  const UnauthorizedWidget({
    Key? key,
    required this.loggedCallback,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.end,
      children: [
        ElevatedButton.icon(
          style: TextButton.styleFrom(
            padding: EdgeInsets.symmetric(
              horizontal: defaultPadding * 1.5,
              vertical: defaultPadding / (Responsive.isMobile(context) ? 2 : 1),
            ),
            backgroundColor: primaryColor,
          ),
          onPressed: () {
            showLoginScreen(context, loggedCallback);
          },
          icon: Icon(Icons.lock_open),
          label: Text("Authorize"),
        ),
      ],
    );
  }
}

class AuthorizedActions extends StatelessWidget {
  final void Function()? update;
  final void Function()? add;
  const AuthorizedActions({
    Key? key,
    this.update,
    this.add,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.end,
      children: [
        ElevatedButton.icon(
          style: TextButton.styleFrom(
            padding: EdgeInsets.symmetric(
              horizontal: defaultPadding * 1.5,
              vertical: defaultPadding / (Responsive.isMobile(context) ? 2 : 1),
            ),
            backgroundColor: primaryColor,
          ),
          onPressed: update,
          icon: Icon(Icons.refresh),
          label: Text("Update"),
        ),
        Container(width: 8),
        ElevatedButton.icon(
          style: TextButton.styleFrom(
            padding: EdgeInsets.symmetric(
              horizontal: defaultPadding * 1.5,
              vertical: defaultPadding / (Responsive.isMobile(context) ? 2 : 1),
            ),
            backgroundColor: primaryColor,
          ),
          onPressed: add,
          icon: Icon(Icons.add),
          label: Text("Add"),
        ),
      ],
    );
  }
}
