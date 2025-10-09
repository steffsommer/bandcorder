export enum EventID {
  RecordingIdle = "RecordingIdle",
  RecordingRunning = "RecordingRunning",
  LiveAudioDataEvent = "LiveAudioData",
}

export interface RunningEventData {
  fileName: string;
  secondsRunning: number;
}
