import 'package:bandcorder/style_constants.dart';
import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';

class ToastService {
  void toastSuccess(String msg) {
    Fluttertoast.showToast(
        msg: msg,
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.TOP,
        backgroundColor: StyleConstants.colorGreen,
        textColor: Colors.black,
        fontSize: StyleConstants.textSizeNormal);
  }

  void toastError(String msg) {
    Fluttertoast.showToast(
        msg: msg,
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.TOP,
        backgroundColor: StyleConstants.colorPurple,
        textColor: Colors.black,
        fontSize: StyleConstants.textSizeNormal);
  }
}
