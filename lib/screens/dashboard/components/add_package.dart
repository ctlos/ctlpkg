import 'package:ctlpkg/generated/v1/pacman.pb.dart';
import 'package:ctlpkg/screens/dashboard/components/button.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../../constants.dart';

showAddPackage(BuildContext context) {
  showDialog(
    context: context,
    barrierColor: Colors.black.withOpacity(0.55),
    builder: (BuildContext context) {
      return AlertDialog(
        backgroundColor: backgroundColor,
        title: AddPackageContents(),
      );
    },
  );
}

class AddPackageContents extends StatefulWidget {
  const AddPackageContents({Key? key}) : super(key: key);

  @override
  State<AddPackageContents> createState() => _AddPackageContentsState();
}

class _AddPackageContentsState extends State<AddPackageContents> {
  var textContoller = TextEditingController();
  String textPlaceholder = "Type AUR packages";
  Widget inonPlaceholder = Icon(
    Icons.archive,
    size: 142,
  );

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisSize: MainAxisSize.min,
      children: [
        Padding(
          padding: const EdgeInsets.all(8.0),
          child: Center(
            child: AnimatedSwitcher(
              duration: Duration(milliseconds: 677),
              child: inonPlaceholder,
            ),
          ),
        ),
        Padding(
          padding: const EdgeInsets.all(8.0),
          child: AnimatedSwitcher(
            duration: Duration(milliseconds: 677),
            child: Text(
              textPlaceholder,
              style: Theme.of(context).textTheme.titleLarge,
            ),
          ),
        ),
        Padding(
          padding: EdgeInsets.all(defaultPadding / 2),
          child: TextFormField(
            controller: textContoller,
            decoration: InputDecoration(
              hintText: "enter packages",
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
          padding: const EdgeInsets.all(defaultPadding),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              CtlButton(
                text: "Close",
                icon: Icons.close,
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
              SizedBox(width: defaultPadding),
              CtlButton(
                icon: Icons.add,
                text: "Download",
                onPressed: () async {
                  setState(() {
                    inonPlaceholder = SpinKitCubeGrid(
                      size: 142,
                      color: Theme.of(context).iconTheme.color,
                    );
                    textPlaceholder = "Processing...";
                  });
                  try {
                    var prefs = await SharedPreferences.getInstance();
                    var token = prefs.getString("token");
                    var resp = await stub.add(AddRequest(
                      packages: textContoller.text.split(" "),
                      token: token,
                    ));
                    setState(() {
                      print(resp);
                      setState(() {
                        inonPlaceholder = Icon(
                          Icons.check,
                          size: 142,
                        );
                        Future.delayed(Duration(milliseconds: 832), () {
                          Navigator.of(context).pop();
                        });
                      });
                    });
                  } catch (e) {
                    setState(() {
                      inonPlaceholder = Icon(
                        Icons.do_disturb,
                        size: 142,
                      );
                      if (e.toString().contains("unable to find package")) {
                        textPlaceholder = "Package not found";
                      } else {
                        textPlaceholder = "Unknown error";
                      }
                    });
                  }
                },
              ),
            ],
          ),
        ),
      ],
    );
  }
}
