import 'package:bandcorder/constants.dart';
import 'package:flutter/material.dart';

class Timer extends StatefulWidget {
  final bool isSpinning;
  final double size;
  final Color color1;
  final Color color2;
  final String centerText;
  final TextStyle? textStyle;

  const Timer({
    super.key,
    required this.isSpinning,
    this.size = 280,
    this.color2 = Constants.colorGreen,
    this.color1 = Constants.colorSurface2,
    this.centerText = '2:49',
    this.textStyle,
  });

  @override
  State<Timer> createState() => _TimerState();
}

class _TimerState extends State<Timer> with SingleTickerProviderStateMixin {
  late AnimationController _controller;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      duration: const Duration(milliseconds: 1000),
      vsync: this,
    );
    if (widget.isSpinning) {
      _controller.repeat();
    }
  }

  @override
  void didUpdateWidget(Timer oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (widget.isSpinning) {
      _controller.repeat();
    } else {
      _controller.stop();
    }
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: widget.size,
      height: widget.size,
      child: Stack(
        alignment: Alignment.center,
        children: [
          AnimatedBuilder(
            animation: _controller,
            builder: (context, child) {
              return Transform.rotate(
                angle: _controller.value * 2 * 3.14159,
                child: CustomPaint(
                  size: Size(widget.size, widget.size),
                  painter: TimerPainter(
                    color1: widget.color1,
                    color2: widget.color2,
                  ),
                ),
              );
            },
          ),
          Text(
            widget.centerText,
            style: const TextStyle(
                fontSize: Constants.textSizeBiggest,
                fontWeight: FontWeight.bold),
            textAlign: TextAlign.center,
          ),
        ],
      ),
    );
  }
}

class TimerPainter extends CustomPainter {
  final Color color1;
  final Color color2;
  static const borderStroke = 2.0;
  static const ringStroke = 50.0;

  TimerPainter({
    required this.color1,
    required this.color2,
  });

  @override
  void paint(Canvas canvas, Size size) {
    final center = Offset(size.width / 2, size.height / 2);
    final radius = size.width / 2;

    final ringRadius = radius - ringStroke / 2 - borderStroke / 2;
    final borderRadius = radius - borderStroke / 2;

    // Draw colored ring
    final paint = Paint()
      ..shader = SweepGradient(
        colors: [
          color1,
          color2,
        ],
        stops: const [0.0, 0.5],
      ).createShader(Rect.fromCircle(center: center, radius: radius))
      ..strokeWidth = ringStroke
      ..style = PaintingStyle.stroke
      ..strokeCap = StrokeCap.round;
    canvas.drawCircle(center, ringRadius, paint);

    // Draw border
    final borderPaint = Paint()
      ..color = Colors.black
      ..strokeWidth = borderStroke
      ..style = PaintingStyle.stroke;
    canvas.drawCircle(center, borderRadius, borderPaint);
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) => false;
}
