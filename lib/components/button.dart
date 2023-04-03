import 'package:ctlpkg/responsive.dart';
import 'package:flutter/material.dart';
import '../../../constants.dart';

class CtlButton extends StatelessWidget {
  final String text;
  final void Function()? onPressed;
  final IconData icon;
  const CtlButton({
    Key? key,
    this.onPressed,
    required this.text,
    required this.icon,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ElevatedButton.icon(
      style: TextButton.styleFrom(
        padding: EdgeInsets.symmetric(
          horizontal: defaultPadding * 1.5,
          vertical: defaultPadding / (Responsive.isMobile(context) ? 2 : 1),
        ),
        backgroundColor: primaryColor,
      ),
      onPressed: onPressed,
      icon: Icon(icon),
      label: Text(text),
    );
  }
}
