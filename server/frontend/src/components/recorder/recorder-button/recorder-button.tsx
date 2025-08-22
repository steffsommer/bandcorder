import { ComponentProps } from "react";
import "./recorder-button.css";

interface Props extends ComponentProps<"button"> {}

export const RecorderButton: React.FC<Props> = ({ className, children }) => {
  return <button className={"button shadow " + (className ?? "")}>{children}</button>;
};
