import { ComponentProps } from "react";
import "./card.css";

export function Card(props: ComponentProps<"div">) {
  return <div className="card">
      {props.children}
  </div>;
}
