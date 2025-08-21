import { Card } from "../card/card";
import { RecorderButton } from "./recorder-button/recorder-button";
import "./recorder.css";
import { Timer } from "./timer/timer";
import { FaFile, FaPause, FaPlay, FaSquareFull } from "react-icons/fa";

export const Recorder: React.FC = () => {
  return (
    <div className="recorder">
      <h2 className="heading">Recorder</h2>
      <Timer isRecording={true} className="timer-widget" onStop={() => { }} />
      <div className="current-file">
        <FaFile />
        <h3>2025-08-21--19-45-00.wav</h3>
      </div>
      <Card className="frequency-card">
        <span>frequency info</span>
      </Card>
      <Card className="volume-card">
        <span>volume info</span>
      </Card>
      <div className="controls">
        <RecorderButton className="play-btn">
          <FaPlay />
        </RecorderButton>
        <RecorderButton className="play-btn">
          <FaPause />
        </RecorderButton>
        <RecorderButton className="play-btn">
          <FaSquareFull />
        </RecorderButton>
      </div>
    </div>
  );
};
