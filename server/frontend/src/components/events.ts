export enum EventID {
  RecordingIdle = "RecordingIdle",
  RecordingRunning = "RecordingRunning",
}

export interface RunningEventData {
  fileName: string;
	secondsRunning: number;
}

