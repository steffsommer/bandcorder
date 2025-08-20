import "./recorder.css";
import { Spinner } from "./spinner/spinner";

export const Recorder: React.FC = () => {
  return (
    <div className="recorder">
      <h2 className="descriptive-header">Recorder</h2>
      <Spinner isRecording={true} />
    </div>
  );
};
