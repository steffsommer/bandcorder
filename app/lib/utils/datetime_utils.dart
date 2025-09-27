/// Return the time difference between now and [timestamp] in format mm:ss
String formatSecondsRunning(int? seconds) {
  if (seconds == null) {
    return "";
  }
  final minutes = (seconds / 60).floor();
  final remainingSeconds = (seconds % 60).floor();
  return '${minutes.toString().padLeft(2, '0')}:${remainingSeconds.toString().padLeft(2, '0')}';
}
