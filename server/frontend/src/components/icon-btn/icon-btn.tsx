import { ComponentProps } from "react";
import "./icon-btn.css";

export function Btn(props: ComponentProps<"button">) {
  return (
    <button className={"btn shadow rounded " + props.className}>
      {props.children}
    </button>
  );
}
