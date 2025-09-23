enum EventId {
  recordingStateUpdate("RecordingState");

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

abstract class Event {
  EventId get id;
}

class RecordingStateEvent extends Event {
  @override
  final EventId id = EventId.recordingStateUpdate;
  final RecordingStateData data;

  RecordingStateEvent(this.data);

  factory RecordingStateEvent.withJsonEventData(
      Map<String, dynamic> eventData) {
    return RecordingStateEvent(RecordingStateData.fromJson(eventData));
  }
}

