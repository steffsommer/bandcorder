import 'dart:convert';

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

class Event {
  late EventId id;

  Event.fromJson(dynamic json) {
    var jsonStr = json.toString();
    var decodedJson = jsonDecode(jsonStr) as Map<String, dynamic>;
    var idStr = decodedJson["eventId"];
    id = EventId.fromString(idStr);
  }
}
