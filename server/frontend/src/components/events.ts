export enum EventID {
  RecordingIdle = "RecordingIdle",
  RecordingRunning = "RecordingRunning",
  LiveAudioDataEvent = "LiveAudioData",
  FileRenamedEvent = "FileRenamed",
  MetronomeIdleEvent = "MetronomeIdle",
  MetronomeRunningEvent = "MetronomeRunning",
}

export interface RunningEventData {
  fileName: string;
  secondsRunning: number;
}
