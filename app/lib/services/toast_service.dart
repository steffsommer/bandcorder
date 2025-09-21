import 'package:bandcorder/constants.dart';
import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';

class ToastService {
  void toastSuccess(String msg) {
    Fluttertoast.showToast(
        msg: msg,
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.TOP,
        backgroundColor: Constants.colorGreen,
        textColor: Colors.black,
        fontSize: Constants.textSizeNormal);
  }

  void toastError(String msg) {
    Fluttertoast.showToast(
        msg: msg,
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.TOP,
        backgroundColor: Constants.colorPurple,
        textColor: Colors.black,
        fontSize: Constants.textSizeNormal);
  }
}
