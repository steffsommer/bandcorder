import { useEffect, useState } from "react";
import "./timer.css";
import "../../../../wailsjs/runtime/runtime";
import { RecordingID as RecordingEvent } from "../../contants";
import { EventsOn } from "../../../../wailsjs/runtime/runtime";
import { TimeUtils } from "../../../utils/time";

interface Props {
  className: string;
}

interface RecordingRunningEvent {
  FileName: string;
  Started: string;
}

export const Timer: React.FC<Props> = ({ className }) => {
  const [durationStr, setDurationStr] = useState("");

  useEffect(() => {
    return EventsOn(RecordingEvent.RUNNING, (data: RecordingRunningEvent) => {
      const str = TimeUtils.SinceStr(data.Started);
      setDurationStr(str);
    });
  });

  return (
    <div className={"timer-container " + (className ? " " + className : "")}>
      <div className={"timer " + (durationStr !== "" ? "active" : "")}></div>
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
