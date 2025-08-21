import { Card } from "../card/card";
import "./recorder.css";
import { Timer } from "./timer/timer";
import { FaFile } from "react-icons/fa";

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
      <div className="controls">controls</div>
    </div>
  );
};
