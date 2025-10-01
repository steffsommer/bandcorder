import { useEffect, useRef } from "react";
import { models } from "../../../wailsjs/go/models";
import { EventsOn } from "../../../wailsjs/runtime/runtime";
import { EventID } from "../events";
import "./volume-bar.css";

const barCount = 40;

export function VolumeBar() {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    drawPercentage(0);
    return EventsOn(
      EventID.LiveAudioDataEvent,
      async (event: models.LiveAudioEventData) => {
        drawPercentage(event.loudnessPercentage);
      },
    );
  }, []);

  function drawPercentage(percentage: number) {
    const context = canvasRef.current?.getContext("2d");
    if (!context) return;

    const width = canvasRef.current?.width ?? 0;
    const height = canvasRef.current?.height ?? 0;
    const barWidth = width / barCount;
    const gap = 2;

    context.clearRect(0, 0, width, height);

    for (let i = 0; i < barCount; i++) {
      const x = i * barWidth;
      const fillThreshold = ((i + 1) / barCount) * 100;
      if (percentage >= fillThreshold) {
        context.fillStyle = "black";
        context.fillRect(x + gap, 0, barWidth - gap * 2, height);
      } else {
        context.strokeStyle = "#aaa";
        context.lineWidth = 1;
        context.strokeRect(x + gap, 0, barWidth - gap * 2, height);
      }
    }
  }

  return (
    <canvas
      className="volume-bar"
      ref={canvasRef}
      width={400}
      height={30}
    ></canvas>
  );
}
