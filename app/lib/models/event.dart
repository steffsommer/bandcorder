enum EventId {
  recordingIdle("RecordingIdle"),
  recordingRunning("RecordingRunning"),
  metronomeStateChange("MetronomeStateChange");

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
  final int secondsRunning;
  final String fileName;

  RecordingRunningEvent({required this.secondsRunning, required this.fileName});

  RecordingRunningEvent.fromJson(Map<String, dynamic> eventData)
      : secondsRunning = eventData['secondsRunning'],
        fileName = eventData['fileName'] ?? '';
}

class RecordingIdleEvent extends Event {}

class MetronomeStateChangeEvent extends Event {
  final bool isRunning;
  final int bpm;

  MetronomeStateChangeEvent({required this.isRunning, required this.bpm});

  MetronomeStateChangeEvent.fromJson(Map<String, dynamic> eventData)
      : isRunning = eventData['isRunning'],
        bpm = eventData['bpm'];
}
