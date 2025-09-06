import { Card } from "../card/card";
import { Button } from "../button/button";
import "./recorder.css";
import { Timer } from "./timer/timer";
import { FaFile, FaPause, FaPlay, FaSquareFull } from "react-icons/fa";
import {
  Abort,
  Start,
  Stop,
} from "../../../wailsjs/go/facades/RecordingFacade.js";

export const Recorder: React.FC = () => {
  return (
    <div className="recorder">
      <h2 className="heading">Recorder</h2>
      <Timer className="timer-widget" />
      <div className="current-file">
        <FaFile size="1.2em" className="file-icon" />
        <h3>2025-08-21--19-45-00.wav</h3>
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
