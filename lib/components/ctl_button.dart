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
          horizontal: 6,
          vertical: 4,
        ),
        backgroundColor: primaryColor,
      ),
      onPressed: onPressed,
      icon: Icon(icon),
      label: Text(text),
    );
  }
}
