export enum EventID {
  RecordingIdle = "RecordingIdle",
  RecordingRunning = "RecordingRunning",
  LiveAudioDataEvent = "LiveAudioData",
  FileRenamedEvent = "FileRenamed",
}

export interface RunningEventData {
  fileName: string;
  secondsRunning: number;
}
