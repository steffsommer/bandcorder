import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';

class ToastService {
  toastSuccess(String message) {
    _toast(message, Colors.black, Colors.green, Toast.LENGTH_SHORT);
  }

  toastError(String message) {
    _toast(message, Colors.white, Colors.red, Toast.LENGTH_LONG);
  }

  _toast(String message, Color textColor, Color backgroundColor, Toast length) {
    Fluttertoast.showToast(
      msg: message,
      toastLength: Toast.LENGTH_SHORT,
      gravity: ToastGravity.TOP,
      backgroundColor: backgroundColor,
      textColor: textColor,
      fontSize: 16.0,
    );
  }
}
