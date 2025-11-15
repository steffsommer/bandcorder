import { AnimatePresence, motion } from "motion/react";
import React, { useEffect, useState } from "react";
import { models } from "../../../wailsjs/go/models";
import { GetRecordings } from "../../../wailsjs/go/facades/FileSystemStorageFacade";
import { EventsOn } from "../../../wailsjs/runtime/runtime";
import { EventID } from "../events";
import { RecordingsListEntry } from "./recordings-list-entry/recordings-list-entry";
import "./recordings-list.css";

export const RecordingsList: React.FC<any> = () => {
  const [recordings, setRecordings] = useState<models.RecordingInfo[]>([]);

  async function updateList() {
    const date = new Date().toISOString();
    const recordingInfos = await GetRecordings(date);
    setRecordings(recordingInfos);
  }

  useEffect(() => {
    let lastEventWasRunningEvent = false;
    updateList();
    const cb1 = EventsOn(EventID.RecordingIdle, async () => {
      if (lastEventWasRunningEvent) {
        updateList();
      }
      lastEventWasRunningEvent = false;
    });
    const cb2 = EventsOn(EventID.RecordingRunning, () => {
      lastEventWasRunningEvent = true;
    });
    const cb3 = EventsOn(EventID.FileRenamedEvent, () => {
      updateList();
    });
    return () => {
      cb1();
      cb2();
      cb3();
    };
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
                initial={{ y: 20, opacity: 0 }}
                animate={{ y: 0, opacity: 1 }}
                transition={{
                  delay: 0.03 * index,
                }}
                exit={{ opacity: 0 }}
                key={item.fileName || `recording-${index}`}
              >
                <RecordingsListEntry recording={item} onChange={updateList} />
              </motion.li>
            ))}
          </AnimatePresence>
        )}
      </div>
    </motion.ul>
  );
};
