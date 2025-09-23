/// Return the time difference between now and [timestamp] in format mm:ss
String formatTimeSince(DateTime timestamp) {
  final now = DateTime.now();
  final difference = now.difference(timestamp);

  final totalSeconds = difference.inSeconds.abs();
  final minutes = totalSeconds ~/ 60;
  final seconds = totalSeconds % 60;

  return '${minutes.toString().padLeft(2, '0')}:${seconds.toString().padLeft(2, '0')}';
}
