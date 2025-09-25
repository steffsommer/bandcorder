import 'package:bandcorder/style_constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

const _marginTop = StyleConstants.spacing;

class CustomAppBar extends StatelessWidget implements PreferredSizeWidget {
  const CustomAppBar({super.key});

  @override
  Widget build(BuildContext context) {
    final canPop = ModalRoute.of(context)?.canPop ?? false;

    return PreferredSize(
      preferredSize: AppBar().preferredSize,
      child: Column(
        children: [
          const SizedBox(height: _marginTop),
          Expanded(
            child: SafeArea(
              child: Container(
                padding: const EdgeInsets.fromLTRB(16, 0, 16, 0),
                color: StyleConstants.colorSurface1,
                child: Container(
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: StyleConstants.borderRadius,
                    border: StyleConstants.border,
                    boxShadow: StyleConstants.boxShadow,
                  ),
                  child: AppBar(
                    shape: RoundedRectangleBorder(
                      borderRadius: StyleConstants.borderRadius,
                    ),
                    leading: canPop ? IconButton(
                      icon: const Icon(Icons.arrow_back, color: Colors.black),
                      onPressed: () => Navigator.of(context).pop(),
                    ) : null,
                    title: Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        SvgPicture.asset("assets/logo.svg",
                            semanticsLabel: "Bandcorder logo", height: 20),
                        const SizedBox(width: 8),
                        const Text(
                          'BANDCORDER',
                          style: TextStyle(color: Colors.black, fontWeight: FontWeight.bold),
                        ),
                      ],
                    ),
                    centerTitle: true,
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
      const Size.fromHeight(kToolbarHeight + _marginTop);
}
