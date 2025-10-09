import 'package:bandcorder/style_constants.dart';
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
        color: StyleConstants.colorSurface2,
        border: StyleConstants.border,
        borderRadius: StyleConstants.borderRadius,
        boxShadow: StyleConstants.boxShadow,
      ),
      padding: StyleConstants.padding,
      width: double.infinity,
      child: child,
    ));
  }
}
