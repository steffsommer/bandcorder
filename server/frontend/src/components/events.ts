export enum EventID {
  RecordingState = "RecordingState",
}

export enum RecordingState {
  RUNNING = "RUNNING",
  IDLE = "IDLE",
}

export interface RecordingStateEvent<T extends RecordingState> {
  State: T;
}

export interface RecordingRunningEvent
  extends RecordingStateEvent<RecordingState.RUNNING> {
  FileName: string;
  Started: string;
}

export type RecordingIdleEvent = RecordingStateEvent<RecordingState.IDLE>;
