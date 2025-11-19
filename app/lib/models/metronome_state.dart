class MetronomeState {
  final bool isRunning;
  final int bpm;

  MetronomeState({required this.isRunning, required this.bpm});

  factory MetronomeState.fromJson(Map<String, dynamic> json) {
    return MetronomeState(
      isRunning: json['isRunning'] ?? false,
      bpm: json['bpm'] ?? 40,
    );
  }
}
