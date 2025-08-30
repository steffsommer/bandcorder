import { useEffect, useState } from "react";
import "./timer.css";
import "../../../../wailsjs/runtime/runtime";
import { RecordingID as RecordingEvent } from "../../contants";
import { EventsOn } from "../../../../wailsjs/runtime/runtime";

interface Props {
  isRecording: boolean;
  className: string;
  onStop: (seconds: number) => void;
}

export const Timer: React.FC<Props> = ({ isRecording, className, onStop }) => {
  const [durationStr, setDurationStr] = useState("");

  useEffect(() => {
    return EventsOn(RecordingEvent.RUNNING, (data) => {
      console.log('Received RUNNING event data!!!!')
      console.log(data)
    })
  });

  useEffect(() => {
    let intervalID: number;

    if (isRecording) {
      const start = Date.now(); // Local variable - immediate value
      const durationStr = getDurationString(start);
      setDurationStr(durationStr);
      intervalID = setInterval(() => {
        const durationStr = getDurationString(start);
        setDurationStr(durationStr);
      }, 1000);
    } else {
      setDurationStr("");
    }

    return () => clearInterval(intervalID);
  }, [isRecording]);
  return (
    <div className={"timer-container " + (className ? " " + className : "")}>
      <div className={"timer " + (isRecording ? "active" : "")}></div>
      <span className="time-label">{durationStr}</span>
    </div>
  );
};

function getDurationString(start: number): string {
  const fullSeconds = Math.floor((Date.now() - start) / 1000);
  const minutesStr = Math.floor(fullSeconds / 60).toString();
  const secondsStr = Math.floor(fullSeconds % 60).toString();
  return `${minutesStr.padStart(2, "0")}:${secondsStr.padStart(2, "0")}`;
}
