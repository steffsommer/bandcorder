import { AnimatePresence, motion } from "motion/react";
import { useEffect, useState } from "react";
import { FaFile, FaPause, FaPlay, FaSquareFull } from "react-icons/fa";
import { FiMic } from "react-icons/fi";
import {
  Abort,
  GetMic,
  Start,
  Stop,
} from "../../../wailsjs/go/facades/RecordingFacade.js";
import { EventsOn } from "../../../wailsjs/runtime/runtime.js";
import { Button } from "../button/button";
import { Card } from "../card/card";
import {
  EventID,
  RecordingRunningEvent,
  RecordingState,
  RecordingStateEvent,
} from "../events.js";
import "./recorder.css";
import { Timer } from "./timer/timer";

export const Recorder: React.FC = () => {
  const [recordingName, setRecordingName] = useState("");
  const [mic, setMic] = useState("");

  const updateMic = async () => {
    try {
      const micStr = await GetMic();
      setMic(micStr);
    } catch (e) {
      console.log("Failed to get microphone: " + e);
    }
  };

  useEffect(() => {
    return EventsOn(EventID.RecordingState, (ev: RecordingStateEvent<any>) => {
      if (ev.state === RecordingState.RUNNING) {
        const runningEvent = ev as RecordingRunningEvent;
        setRecordingName(runningEvent.fileName);
      } else {
        setRecordingName("");
      }
    });
  });

  useEffect(() => {
    updateMic();
    return EventsOn(EventID.SettingsChanged, async () => {
      await updateMic();
    });
  });

  return (
    <div className="recorder">
      <h2 className="heading">Recorder</h2>
      <Timer className="timer-widget" />
      <div className="infos">
        <AnimatePresence>
          {recordingName && (
            <motion.div
              initial={{ y: 50, opacity: 0 }}
              animate={{ y: 0, opacity: 1 }}
              exit={{ opacity: 0 }}
              className="current-file-info"
            >
              <FaFile size="1.2em" className="file-icon" />
              <h3>{recordingName}</h3>
            </motion.div>
          )}
        </AnimatePresence>
        <div className="mic-info">
          <FiMic size="1.3em" />
          <h3>{mic}</h3>
        </div>
      </div>
      <Card className="frequency-card">
        <span>🚧 frequency info 🚧</span>
      </Card>
      <Card className="volume-card">
        <span>🚧 volume info 🚧</span>
      </Card>
      <div className="controls">
        <Button onClick={Start} className="recorder-btn icon-large play-btn">
          <FaPlay />
        </Button>
        <Button onClick={Stop} className="recorder-btn icon-large pause-btn">
          <FaPause />
        </Button>
        <Button onClick={Abort} className="recorder-btn icon-large abort-btn">
          <FaSquareFull />
        </Button>
      </div>
    </div>
  );
};
