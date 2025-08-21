import { ComponentProps } from "react";
import "./recorder-button.css";

export const RecorderButton: React.FC<ComponentProps<"button">> = ({
  className,
  children,
}) => {
  return (
    <button className={"recorder-btn shadow " + (className ?? "")}>{children}</button>
  );
};
