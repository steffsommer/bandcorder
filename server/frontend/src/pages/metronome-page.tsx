import styles from "./metronome-page.module.css";
import { Start, Stop, UpdateBpm } from "../../wailsjs/go/services/MetronomeService";

export default function Metronome() {
  const isPlaying = false;
  const bpm = 120;
  function toggle() {
    console.log("not yet implemented");
    Start();
  }

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <div className={styles.barsContainer}>
          {[...Array(8)].map((_, i) => (
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
          className={`${styles.button} ${isPlaying ? styles.buttonStop : styles.buttonStart}`}
        >
          {isPlaying ? "STOP" : "START"}
        </button>
      </div>
    </div>
  );
}
