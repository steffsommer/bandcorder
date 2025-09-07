import React, { useEffect, useRef, useState } from "react";
import { models } from "../../../wailsjs/go/models";
import { GetRecordings } from "../../../wailsjs/go/services/FileSystemStorageService";
import { EventsOn } from "../../../wailsjs/runtime/runtime";
import { EventID, RecordingState, RecordingStateEvent } from "../events";
import { RecordingsListEntry } from "./recordings-list-entry/recordings-list-entry";
import "./recordings-list.css";

export const RecordingsList: React.FC<any> = () => {
  const [recordings, setRecordings] = useState<models.RecordingInfo[]>([]);
  const lastState = useRef<RecordingState | null>(null);

  useEffect(() => {
    return EventsOn(
      EventID.RecordingState,
      async (ev: RecordingStateEvent<any>) => {
        if (
          ev.State === RecordingState.IDLE &&
          lastState.current !== RecordingState.IDLE
        ) {
          const date = new Date().toISOString();
          const recordingInfos = await GetRecordings(date);
          setRecordings(recordingInfos);
        }
        lastState.current = ev.State;
      },
    );
  }, []);
  return (
    <div className="recordings-list">
      <h2 className="descriptive-header">Todays' recordings</h2>
      <div className="recordings">
        {recordings.length === 0 ? (
          <h2 className="no-items">No items to display</h2>
        ) : (
          recordings.map((item, index) => (
            <RecordingsListEntry recording={item} key={index} />
          ))
        )}
      </div>
    </div>
  );
};
