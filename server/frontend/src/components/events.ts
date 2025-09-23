export enum EventID {
  RecordingState = "RecordingState",
  SettingsChanged = "SettingsUpdate",
}

export enum RecordingState {
  RUNNING = "RUNNING",
  IDLE = "IDLE",
}

export interface RecordingStateEvent<T extends RecordingState> {
  state: T;
}

export interface RecordingRunningEvent
  extends RecordingStateEvent<RecordingState.RUNNING> {
  fileName: string;
  started: string;
}

export type RecordingIdleEvent = RecordingStateEvent<RecordingState.IDLE>;
