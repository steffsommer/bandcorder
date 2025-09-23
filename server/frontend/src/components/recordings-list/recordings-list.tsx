import React, { useEffect, useRef, useState } from "react";
import { models } from "../../../wailsjs/go/models";
import { GetRecordings } from "../../../wailsjs/go/services/FileSystemStorageService";
import { EventsOn } from "../../../wailsjs/runtime/runtime";
import { EventID, RecordingState, RecordingStateEvent } from "../events";
import { RecordingsListEntry } from "./recordings-list-entry/recordings-list-entry";
import "./recordings-list.css";
import { AnimatePresence, motion } from "motion/react";

export const RecordingsList: React.FC<any> = () => {
  const [recordings, setRecordings] = useState<models.RecordingInfo[]>([]);
  const lastState = useRef<RecordingState | null>(null);

  useEffect(() => {
    return EventsOn(
      EventID.RecordingState,
      async (ev: RecordingStateEvent<any>) => {
        if (
          ev.state === RecordingState.IDLE &&
          lastState.current !== RecordingState.IDLE
        ) {
          const date = new Date().toISOString();
          const recordingInfos = await GetRecordings(date);
          setRecordings(recordingInfos);
        }
        lastState.current = ev.state;
      },
    );
  }, []);

  return (
    <motion.ul layout layoutId={"list"} className="recordings-list">
      <h2 className="descriptive-header">Todays' recordings</h2>
      <div className="recordings">
        {recordings.length === 0 ? (
          <h2 className="no-items">No items to display</h2>
        ) : (
          <AnimatePresence>
            {recordings.map((item, index) => (
              <motion.li
                initial={{ y: 50, opacity: 0 }}
                animate={{ y: 0, opacity: 1 }}
                transition={{
                  delay: 0.03 * index,
                }}
                exit={{ opacity: 0 }}
                key={item.FileName || `recording-${index}`}
              >
                <RecordingsListEntry recording={item} />
              </motion.li>
            ))}
          </AnimatePresence>
        )}
      </div>
    </motion.ul>
  );
};
