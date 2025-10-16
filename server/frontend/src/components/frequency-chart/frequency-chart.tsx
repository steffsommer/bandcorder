import { useRef, useEffect } from "react";
import { models } from "../../../wailsjs/go/models";
import { EventsOn } from "../../../wailsjs/runtime/runtime";
import { EventID } from "../events";
import "./frequency-chart.css";

const BAR_COUNT = 40;
const CANVAS_WIDTH = 1000;
const CANVAS_HEIGHT = 400;
const MIN_BAR_HEIGHT = 2;
const BAR_MIN = 3;
const DEFAULT_BARS = Array(40).fill(BAR_MIN);

export function FrequencyChart() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const animationFrameHandleRef = useRef(0);
  let bars: number[] = DEFAULT_BARS;

  useEffect(() => {
    requestAnimationFrame(draw);
    const unsubAudioData = EventsOn(EventID.LiveAudioDataEvent,
      async (event: models.LiveAudioEventData) => {
        // interpolator.setTarget(event.loudnessPercentage);
        bars = event.frequencyBars.map(val => val < BAR_MIN ? BAR_MIN : val)
      }
    );
    const unsubIdleEvent = EventsOn(EventID.RecordingIdle,
      async () => {
        bars = DEFAULT_BARS;
      }
    );
    return () => {
      unsubAudioData();
      unsubIdleEvent();
      cancelAnimationFrame(animationFrameHandleRef.current);
    };
  }, []);

  function draw() {
    if (!canvasRef.current) {
      throw new Error("Failed to draw. Canvas Ref is falsy")
    }
    const context = canvasRef.current?.getContext("2d");
    if (!context) {
      throw new Error("Failed to draw. Canvas Context was falsy")
    }
    context.clearRect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT);
    const canvas = canvasRef.current;
    if (!canvas) return;
    const ctx = canvas.getContext("2d");
    if (!ctx) return;
    ctx.clearRect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT);

    const barWidth = CANVAS_WIDTH / BAR_COUNT;
    bars.slice(0, BAR_COUNT).forEach((value, i) => {
      const barHeight =
        ((CANVAS_HEIGHT - MIN_BAR_HEIGHT) * value) / 100 + MIN_BAR_HEIGHT;
      ctx.fillStyle = "#3498db";
      ctx.fillRect(
        i * barWidth,
        CANVAS_HEIGHT - barHeight,
        barWidth - 2,
        barHeight
      );
    });
    requestAnimationFrame(draw);
  }

  // Initial blank draw
  useEffect(() => {
    const canvas = canvasRef.current;
    if (canvas) {
      const ctx = canvas.getContext("2d");
      if (ctx) ctx.clearRect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT);
    }
  }, []);

  // Attach draw to the component for external access
  (FrequencyChart as any).draw = draw;

  return (
    <canvas
      className="frequency-chart"
      ref={canvasRef}
      width={CANVAS_WIDTH}
      height={CANVAS_HEIGHT}
      style={{ border: "1px solid #ccc" }}
    />
  );
}
