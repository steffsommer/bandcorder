import 'package:bandcorder/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

const _marginTop = Constants.spacing;

class CustomAppBar extends StatelessWidget implements PreferredSizeWidget {
  const CustomAppBar({super.key});

  @override
  Widget build(BuildContext context) {
    return PreferredSize(
      preferredSize: AppBar().preferredSize,
      child: Column(
        children: [
          const SizedBox(height: _marginTop), // 20px margin at top
          Expanded(
            child: SafeArea(
              child: Container(
                padding: const EdgeInsets.fromLTRB(16, 0, 16, 0),
                color: Constants.colorSurface1,
                child: Container(
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: Constants.borderRadius,
                    border: Constants.border,
                    boxShadow: Constants.boxShadow,
                  ),
                  child: AppBar(
                    shape: RoundedRectangleBorder(
                      borderRadius: Constants.borderRadius,
                    ),
                    title: Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        SvgPicture.asset("assets/logo.svg",
                            semanticsLabel: "Bandcorder logo", height: 26),
                        const SizedBox(width: 8),
                        const Text(
                          'BANDCORDER',
                          style: TextStyle(color: Colors.black),
                        ),
                      ],
                    ),
                    elevation: 0,
                    backgroundColor: Colors.transparent,
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  @override
  Size get preferredSize =>
      const Size.fromHeight(kToolbarHeight + _marginTop); // Increased height
}
