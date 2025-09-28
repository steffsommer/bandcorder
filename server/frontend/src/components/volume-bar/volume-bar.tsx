import { useEffect, useRef } from "react";
import "./volume-bar.css";

export function VolumeBar() {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    const context = canvasRef.current?.getContext("2d");
    if (!context) {
      return;
    }
    const width = canvasRef.current?.width ?? 0;
    const height = (canvasRef.current?.height ?? 0);
    context.fillStyle = "black";
    context.rect(0, 0, width, height);
    context.fill();
  }, [canvasRef.current]);

  return <canvas className="volume-bar" ref={canvasRef}></canvas>;
}
