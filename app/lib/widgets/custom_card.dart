import 'package:bandcorder/contants.dart';
import 'package:flutter/cupertino.dart';

class CustomCard extends StatelessWidget {
  final Widget? child;

  const CustomCard({super.key, this.child});

  @override
  Widget build(
    BuildContext context,
  ) {
    return IntrinsicHeight(
        child: Container(
      decoration: BoxDecoration(
        color: Constants.colorSurface2,
        border: Constants.border,
        borderRadius: Constants.borderRadius,
        boxShadow: Constants.boxShadow,
      ),
      padding: Constants.padding,
      child: child,
    ));
  }
}
