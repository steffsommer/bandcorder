import 'package:flutter/material.dart';

class LoadingOverlayService {
  static OverlayEntry? _overlayEntry;
  static bool _isVisible = false;
  static GlobalKey<NavigatorState>? _navigatorKey;

  static void init(GlobalKey<NavigatorState> navigatorKey) {
    _navigatorKey = navigatorKey;
  }

  static void showLoading({String text = 'Loading...'}) {
    if (_isVisible) return;
    if (_navigatorKey?.currentContext == null) {
      debugPrint('LoadingService: No valid context found');
      return;
    }

    // Ensure we run this after the frame is built
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final context = _navigatorKey!.currentContext!;
      
      // Check if there's an Overlay widget in the tree
      if (Overlay.of(context, rootOverlay: true) == null) {
        debugPrint('LoadingService: No Overlay widget found');
        return;
      }

      _overlayEntry = OverlayEntry(
        builder: (context) => Material(
          color: Colors.transparent,
          child: Stack(
            children: [
              Container(
                color: Colors.black.withOpacity(0.5),
              ),
              Center(
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    const CircularProgressIndicator(
                      valueColor: AlwaysStoppedAnimation<Color>(Colors.white),
                    ),
                    const SizedBox(height: 16),
                    Text(
                      text,
                      style: const TextStyle(
                        color: Colors.white,
                        fontSize: 16,
                        decoration: TextDecoration.none,
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      );

      Overlay.of(context, rootOverlay: true).insert(_overlayEntry!);
      _isVisible = true;
    });
  }

  static void hideLoading() {
    if (!_isVisible) return;
    
    _overlayEntry?.remove();
    _overlayEntry = null;
    _isVisible = false;
  }
}