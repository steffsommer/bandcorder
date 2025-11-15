import { useEffect, useState } from "react";
import { Start, Stop, UpdateBpm } from "../../wailsjs/go/services/MetronomeService";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { EventID } from "../components/events";
import styles from "./metronome-page.module.css";

export default function Metronome() {
  const bpm = 120;
  const barCount = 8;

  const [on, setOn] = useState(false);

  function toggle() {
    if (on) {
      Stop();
    } else {
      Start();
    }
  }

  useEffect(() => {
    const cb1 = EventsOn(EventID.MetronomeRunningEvent, () => {
      setOn(true);
    });
    const cb2 = EventsOn(EventID.MetronomeIdleEvent, () => {
      setOn(false);
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
            <div key={i} className={`${styles.bar} ${i === 0 ? styles.barActive : ""}`} />
          ))}
        </div>

        <div className={styles.display}>
          <div className={styles.bpmValue}>80</div>
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
