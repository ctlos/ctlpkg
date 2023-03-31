import 'package:flutter/material.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';

import '../../../constants.dart';

showUpdateNotification(BuildContext context) {
  showBottomSheet(
    context: context,
    builder: (context) {
      return Text("Updating packages...");
    },
  );
}

class UpdateNotification extends StatelessWidget {
  const UpdateNotification({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 32,
      child: Row(
        mainAxisSize: MainAxisSize.max,
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Icon(
            Icons.update,
            size: 24,
          ),
          SizedBox(width: defaultPadding),
          Text(
            "Updating packages...",
            style: Theme.of(context).textTheme.titleLarge,
          ),
          SizedBox(width: defaultPadding),
          SpinKitCubeGrid(
            size: 24,
            color: Theme.of(context).iconTheme.color,
          ),
        ],
      ),
    );
  }
}
