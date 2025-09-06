import { useEffect, useState } from "react";
import "./timer.css";
import "../../../../wailsjs/runtime/runtime";
import {
  EventID,
  RecordingRunningEvent,
  RecordingState,
  RecordingStateEvent,
} from "../../events";
import { EventsOn } from "../../../../wailsjs/runtime/runtime";
import { TimeUtils } from "../../../utils/time";

interface Props {
  className: string;
}

export const Timer: React.FC<Props> = ({ className }) => {
  const [durationStr, setDurationStr] = useState("");

  useEffect(() => {
    return EventsOn(EventID.RecordingState, (ev: RecordingStateEvent<any>) => {
      if (ev.State === RecordingState.RUNNING) {
        const runningEvent = ev as RecordingRunningEvent;
        const str = TimeUtils.SinceStr(runningEvent.Started);
        setDurationStr(str);
      } else {
        setDurationStr("");
      }
    });
  });

  return (
    <div className={"timer-container " + (className ? " " + className : "")}>
      <div className={"timer " + (durationStr !== "" ? "active" : "")}></div>
      <span className="time-label">{durationStr}</span>
    </div>
  );
};
