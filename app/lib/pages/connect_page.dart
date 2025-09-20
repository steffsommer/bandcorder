import 'package:bandcorder/contants.dart';
import 'package:bandcorder/widgets/custom_app_bar.dart';
import 'package:bandcorder/pages/record_page.dart';
import 'package:bandcorder/widgets/custom_card.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import '../services/web_socket_service.dart';
import '../widgets/custom_text_field.dart';
import '../widgets/custom_button.dart';

class ConnectPage extends StatefulWidget {
  const ConnectPage({super.key});

  @override
  ConnectPageState createState() => ConnectPageState();
}

class ConnectPageState extends State<ConnectPage> {
  final WebSocketService _socketService = WebSocketService();

  String _textFieldValue = '10.0.2.2';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppBar(),
      body: Padding(
        padding: Constants.padding,
        child: CustomCard(
          child: Column(
            children: [
              SvgPicture.asset("assets/desktop_pc.svg",
                  semanticsLabel: "Bandcorder logo", height: 160),
              const SizedBox(height: 30),
              CustomTextField(
                defaultValue: _textFieldValue,
                labelText:
                    'Please enter the IP displayed in the Bandcorder Desktop application',
                onChanged: (value) {
                  setState(() {
                    _textFieldValue = value;
                  });
                },
              ),
              const SizedBox(height: 30),
              CustomButton(
                color: Constants.colorGreen,
                children: const [
                  Text("Connect"),
                ],
                onPressed: () async {
                  await _socketService.connect(_textFieldValue);
                  Navigator.push(
                    context,
                    MaterialPageRoute<void>(
                      builder: (context) => const RecordPage(),
                    ),
                  );
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}
