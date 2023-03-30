import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:ctlpkg/responsive.dart';
import 'package:flutter/material.dart';

import '../../../constants.dart';

showLoginScreen(BuildContext context) {
  showDialog(
    context: context,
    barrierColor: secondaryColor,
    builder: (BuildContext context) {
      return AlertDialog(
        backgroundColor: backgroundColor,
        title: LoginContents(),
      );
    },
  );
}

class LoginContents extends StatelessWidget {
  final loginController = TextEditingController();
  final passwordController = TextEditingController();

  LoginContents({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisSize: MainAxisSize.min,
      children: [
        Center(
          child: Padding(
            padding: const EdgeInsets.all(8.0),
            child: Icon(
              Icons.lock,
              size: 92,
            ),
          ),
        ),
        Padding(
          padding: EdgeInsets.all(defaultPadding / 2),
          child: TextFormField(
            controller: loginController,
            decoration: InputDecoration(
              hintText: "login",
              fillColor: secondaryColor,
              filled: true,
              border: OutlineInputBorder(
                borderSide: BorderSide.none,
                borderRadius: const BorderRadius.all(Radius.circular(10)),
              ),
            ),
          ),
        ),
        Padding(
          padding: EdgeInsets.all(defaultPadding / 2),
          child: TextFormField(
            obscureText: true,
            controller: passwordController,
            decoration: InputDecoration(
              hintText: "password",
              fillColor: secondaryColor,
              filled: true,
              border: OutlineInputBorder(
                borderSide: BorderSide.none,
                borderRadius: const BorderRadius.all(Radius.circular(10)),
              ),
            ),
          ),
        ),
        Center(
          child: Padding(
            padding: const EdgeInsets.all(defaultPadding / 2),
            child: ElevatedButton.icon(
              style: TextButton.styleFrom(
                padding: EdgeInsets.symmetric(
                  horizontal: defaultPadding * 1.5,
                  vertical:
                      defaultPadding / (Responsive.isMobile(context) ? 2 : 1),
                ),
                backgroundColor: primaryColor,
              ),
              onPressed: () async {
                try {
                  var resp = await stub.login(LoginRequest(
                    login: loginController.text,
                    password: passwordController.text,
                  ));
                  storage.setItem("token", resp.token);
                  Navigator.of(context).pop();
                } catch (e) {
                  print("failed $e");
                }
              },
              icon: Icon(Icons.person),
              label: Text("Authorize"),
            ),
          ),
        ),
      ],
    );
  }
}
