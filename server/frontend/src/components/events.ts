export enum EventID {
  RecordingIdle = "RecordingIdle",
  RecordingRunning = "RecordingRunning",
  LiveAudioDataEvent = "LiveAudioData",
  FileRenamedEvent = "FileRenamed",
  MetronomeStateChangeEvent = "MetronomeStateChange",
  MetronomeBeatEvent = "MetronomeBeat",
}

export interface RunningEventData {
  fileName: string;
  secondsRunning: number;
}
