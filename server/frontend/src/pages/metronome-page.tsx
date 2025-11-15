import { useEffect, useState } from "react";
import { Start, Stop, UpdateBpm } from "../../wailsjs/go/services/MetronomeService";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { EventID, MetronomeRunningData } from "../components/events";
import styles from "./metronome-page.module.css";

export default function Metronome() {
  const [bpm, setBpm] = useState(120);
  const barCount = 4;
  const [on, setOn] = useState(false);
  const [activeBar, setActiveBar] = useState(-1);

  function toggle() {
    if (on) {
      Stop();
    } else {
      Start();
    }
  }

  useEffect(() => {
    const cb1 = EventsOn(EventID.MetronomeRunningEvent, (data: MetronomeRunningData) => {
      setOn(true);
      setActiveBar(data.beatCount % barCount);
    });
    const cb2 = EventsOn(EventID.MetronomeIdleEvent, () => {
      setOn(false);
      setActiveBar(-1);
    });
    return () => {
      cb1();
      cb2();
    };
  }, []);

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <div className={styles.barsContainer}>
          {[...Array(barCount)].map((_, i) => (
            <div key={i} className={`${styles.bar} ${i === activeBar ? styles.barActive : ""}`} />
          ))}
        </div>
        <div className={styles.display}>
          <div className={styles.bpmValue}>{bpm}</div>
        </div>
        <div className={styles.sliderContainer}>
          <input
            type="range"
            value={bpm}
            min="40"
            max="240"
            className={styles.slider}
            style={{
              background: `linear-gradient(to right, #FCD34D ${((bpm - 40) / 200) * 100}%, white ${((bpm - 40) / 200) * 100}%)`,
            }}
            onChange={(e) => setBpm(Number(e.target.value))}
            onMouseUp={(e) => UpdateBpm(Number((e.target as HTMLInputElement).value))}
            onTouchEnd={(e) => UpdateBpm(Number((e.target as HTMLInputElement).value))}
          />
        </div>
        <button
          onClick={toggle}
          className={`${styles.button} ${on ? styles.buttonStop : styles.buttonStart}`}
        >
          {on ? "STOP" : "START"}
        </button>
      </div>
    </div>
  );
}
