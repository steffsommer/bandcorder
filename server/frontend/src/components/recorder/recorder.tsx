import { Card } from "../card/card";
import "./recorder.css";
import { Spinner } from "./spinner/spinner";

export const Recorder: React.FC = () => {
  return (
    <div className="recorder">
      <h2 className="descriptive-header">Recorder</h2>
      <Spinner isRecording={true} className="spinner" />
      <div className="current-file">
        <span>current file info</span>
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
