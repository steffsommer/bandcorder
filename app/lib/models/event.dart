enum EventId {
  recordingIdle("RecordingIdle"),
  recordingRunning("RecordingRunning");

  const EventId(this.value);

  final String value;

  static EventId fromString(String value) {
    try {
      return EventId.values.firstWhere((e) => e.value == value);
    } catch (e) {
      throw ArgumentError("Value '$value' is not a valid Recording Event ID");
    }
  }
}

// Data types for different events
class RecordingStateData {
  final bool isRecording;
  final String fileName;

  RecordingStateData({required this.isRecording, required this.fileName});

  factory RecordingStateData.fromJson(Map<String, dynamic> json) {
    return RecordingStateData(
      isRecording: json['isRecording'] ?? false,
      fileName: json['fileName'] ?? '',
    );
  }
}

// Marker
abstract class Event {}

class RecordingRunningEvent extends Event {
  final DateTime started;
  final String fileName;

  RecordingRunningEvent.fromJson(Map<String, dynamic> eventData)
      : started = DateTime.parse(eventData['started']),
        fileName = eventData['fileName'] ?? '';
}

class RecordingIdleEvent extends Event {}
