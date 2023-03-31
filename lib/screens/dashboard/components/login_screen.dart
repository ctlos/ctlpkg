import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:ctlpkg/responsive.dart';
import 'package:flutter/material.dart';
import 'package:flutter_shake_animated/flutter_shake_animated.dart';
import 'package:shared_preferences/shared_preferences.dart';
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

class LoginContents extends StatefulWidget {
  LoginContents({Key? key}) : super(key: key);

  @override
  State<LoginContents> createState() => _LoginContentsState();
}

class _LoginContentsState extends State<LoginContents> {
  final loginController = TextEditingController();
  final passwordController = TextEditingController();
  var shaking = false;

  @override
  Widget build(BuildContext context) {
    return ShakeWidget(
      shakeConstant: ShakeDefaultConstant1(),
      autoPlay: shaking,
      enableWebMouseHover: false,
      child: Column(
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
                    var prefs = await SharedPreferences.getInstance();
                    prefs.setString("token", resp.token);
                    Navigator.of(context).pop();
                  } catch (e) {
                    setState(() {
                      shaking = true;
                    });
                    await Future.delayed(Duration(milliseconds: 227), () {
                      setState(() {
                        shaking = false;
                      });
                    });
                  }
                },
                icon: Icon(Icons.person),
                label: Text("Authorize"),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
