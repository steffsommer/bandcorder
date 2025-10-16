import { useEffect, useRef } from "react";
import { models } from "../../../wailsjs/go/models";
import { EventsOn } from "../../../wailsjs/runtime/runtime";
import { LinearInterpolator } from "../../utils/linear-interpolator";
import { EventID } from "../events";
import "./volume-bar.css";

const BAR_COUNT = 40;
const INTERPOLATION_SPEED = 60;

export function VolumeBar() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const interpolator = new LinearInterpolator(INTERPOLATION_SPEED);
  const animationFrameHandleRef = useRef(0);

  useEffect(() => {
    requestAnimationFrame(draw);
    const unsub = EventsOn(EventID.LiveAudioDataEvent, async (event: models.LiveAudioEventData) => {
      interpolator.setTarget(event.loudnessPercentage);
    });

    return () => {
      unsub();
      cancelAnimationFrame(animationFrameHandleRef.current);
    };
  }, []);

  function draw() {
    const percentage = interpolator.getValue();
    const context = canvasRef.current?.getContext("2d");
    if (!context) return;

    const width = canvasRef.current?.width ?? 0;
    const height = canvasRef.current?.height ?? 0;
    const barWidth = width / BAR_COUNT;
    const gap = 2;

    context.clearRect(0, 0, width, height);

    for (let i = 0; i < BAR_COUNT; i++) {
      const x = i * barWidth;
      const fillThreshold = ((i + 1) / BAR_COUNT) * 100;
      if (percentage >= fillThreshold) {
        context.fillStyle = "black";
        context.fillRect(x + gap, 0, barWidth - gap * 2, height);
      } else {
        context.strokeStyle = "#aaa";
        context.lineWidth = 1;
        context.strokeRect(x + gap, 0, barWidth - gap * 2, height);
      }
    }
    animationFrameHandleRef.current = requestAnimationFrame(draw);
  }

  return <canvas className="volume-bar" ref={canvasRef} width={1000} height={120}></canvas>;
}
