import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:flutter/material.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';
import 'package:grpc/grpc_web.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../constants.dart';

showUpdateNotification(BuildContext context) {
  showBottomSheet(
    context: context,
    builder: (context) {
      return UpdateNotification();
    },
  );
}

class UpdateNotification extends StatefulWidget {
  const UpdateNotification({Key? key}) : super(key: key);

  @override
  State<UpdateNotification> createState() => _UpdateNotificationState();
}

class _UpdateNotificationState extends State<UpdateNotification> {
  Widget content = Placeholder();

  dropAfterUpdate(BuildContext context) async {
    var prefs = await SharedPreferences.getInstance();
    try {
      var resp = await stub.update(UpdateRequest(
        token: prefs.getString("token"),
      ));
      setState(() {
        content = Row(
          key: UniqueKey(),
          mainAxisSize: MainAxisSize.max,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(Icons.check, size: 24),
          ],
        );
        Future.delayed(Duration(milliseconds: 932), () {
          Navigator.of(context).pop();
        });
      });
    } on GrpcError catch (e) {
      setState(() {
        content = Row(
          key: UniqueKey(),
          mainAxisSize: MainAxisSize.max,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(Icons.do_disturb, size: 24),
            SizedBox(width: defaultPadding),
            Text(
              e.message ?? "Unknown gRPC error",
              style: Theme.of(context).textTheme.titleLarge,
            ),
          ],
        );
        Future.delayed(Duration(seconds: 15), () {
          Navigator.of(context).pop();
        });
      });
    } catch (e) {
      print("unexpected error");
    }
  }

  @override
  void initState() {
    content = Row(
      key: UniqueKey(),
      mainAxisSize: MainAxisSize.max,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        SpinKitDoubleBounce(
          size: 24,
          color: Theme.of(context).iconTheme.color,
        ),
        SizedBox(width: defaultPadding),
        Text(
          "Updating packages...",
          style: Theme.of(context).textTheme.titleLarge,
        ),
      ],
    );
    super.initState();
    dropAfterUpdate(context);
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 72,
      decoration: BoxDecoration(
        color: secondaryColor,
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.35),
            blurRadius: 1,
            offset: const Offset(-1, -1),
          ),
        ],
      ),
      child: Padding(
        padding: const EdgeInsets.all(defaultPadding),
        child: AnimatedSwitcher(
          duration: Duration(milliseconds: 458),
          child: content,
        ),
      ),
    );
  }
}
