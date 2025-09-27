import { useEffect, useState } from "react";
import "../../../../wailsjs/runtime/runtime";
import { EventsOn } from "../../../../wailsjs/runtime/runtime";
import { TimeUtils } from "../../../utils/time";
import { EventID, RunningEventData } from "../../events";
import "./timer.css";

interface Props {
  className: string;
}

export const Timer: React.FC<Props> = ({ className }) => {
  const [durationStr, setDurationStr] = useState("");

  useEffect(() => {
    const cb1 = EventsOn(EventID.RecordingIdle, () => {
      setDurationStr("");
    });
    const cb2 = EventsOn(EventID.RecordingRunning, (data: RunningEventData) => {
      const str = TimeUtils.toMMSS(data.secondsRunning);
      setDurationStr(str);
    });
    return () => {
      cb1();
      cb2();
    };
  }, []);

  return (
    <div className={"timer-container " + (className ? " " + className : "")}>
      <div
        className={"gradient-spinner " + (durationStr !== "" ? "active" : "")}
      ></div>
      <span className="duration-label">{durationStr}</span>
    </div>
  );
};
