import { AnimatePresence, motion } from "motion/react";
import { useEffect, useState } from "react";
import { FaFile, FaPause, FaPlay, FaSquareFull } from "react-icons/fa";
import { FiMic } from "react-icons/fi";
import { Abort, GetMic, Start, Stop } from "../../../wailsjs/go/services/RecordingService";
import { EventsOn } from "../../../wailsjs/runtime/runtime.js";
import { toastFailure } from "../../services/toast-service/toast-service";
import { Button } from "../button/button";
import { Card } from "../card/card";
import { EventID, RunningEventData } from "../events.js";
import { FrequencyChart } from "../frequency-chart/frequency-chart.js";
import { VolumeBar } from "../volume-bar/volume-bar.js";
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

  const start = async () => {
    try {
      await Start();
    } catch (e) {
      toastFailure(String(e));
    }
  };

  const stop = async () => {
    try {
      await Stop();
    } catch (e) {
      toastFailure(String(e));
    }
  };

  const abort = async () => {
    try {
      await Abort();
    } catch (e) {
      toastFailure(String(e));
    }
  };

  useEffect(() => {
    const cb1 = EventsOn(EventID.RecordingIdle, () => {
      setRecordingName("");
    });
    const cb2 = EventsOn(EventID.RecordingRunning, (data: RunningEventData) => {
      setRecordingName(data.fileName);
    });
    return () => {
      cb1();
      cb2();
    };
  });

  useEffect(() => {
    updateMic();
  });

  return (
    <div className="recorder">
      <h2 className="heading">Recorder</h2>
      <Timer className="timer-widget" />
      <div className="infos">
        <AnimatePresence>
          {recordingName && (
            <motion.div
              initial={{ y: 20, opacity: 0 }}
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
        <FrequencyChart />
      </Card>
      <Card className="volume-card">
        <VolumeBar />
      </Card>
      <div className="controls">
        <Button onClick={start} className="recorder-btn icon-large play-btn">
          <FaPlay />
        </Button>
        <Button onClick={stop} className="recorder-btn icon-large pause-btn">
          <FaPause />
        </Button>
        <Button onClick={abort} className="recorder-btn icon-large abort-btn">
          <FaSquareFull />
        </Button>
      </div>
    </div>
  );
};
