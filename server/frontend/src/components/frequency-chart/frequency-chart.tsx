import { useRef, useEffect } from "react";
import { models } from "../../../wailsjs/go/models";
import { EventsOn } from "../../../wailsjs/runtime/runtime";
import { LinearInterpolator } from "../../utils/linear-interpolator";
import { EventID } from "../events";
import "./frequency-chart.css";

const BAR_COUNT = 40;
const MIN_BAR_HEIGHT = 3;
const CANVAS_WIDTH = 1000;
const CANVAS_HEIGHT = 400;
const INTERPOLATION_SPEED = 200;

export function FrequencyChart() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const animationFrameHandleRef = useRef(0);
  let interpolators: LinearInterpolator[] = [];

  useEffect(() => {
    for (let i = 0; i < BAR_COUNT; i++) {
      interpolators.push(new LinearInterpolator(INTERPOLATION_SPEED));
    }
    requestAnimationFrame(draw);
    const unsubAudioData = EventsOn(EventID.LiveAudioDataEvent,
      async (event: models.LiveAudioEventData) => {
        const bars = event.frequencyBars.map(val => val < MIN_BAR_HEIGHT ? MIN_BAR_HEIGHT : val)
        for (let i = 0; i < interpolators.length; i++) {
          interpolators[i].setTarget(bars[i]);
        }
      }
    );
    const unsubIdleEvent = EventsOn(EventID.RecordingIdle,
      async () => {
        interpolators.forEach(ip => ip.setTarget(0));
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
    const ctx = canvasRef.current?.getContext("2d");
    if (!ctx) {
      throw new Error("Failed to draw. Canvas Context was falsy")
    }
    ctx.clearRect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT);

    const barWidth = CANVAS_WIDTH / BAR_COUNT;
    const bars = interpolators.map(ip => ip.getValue())
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
