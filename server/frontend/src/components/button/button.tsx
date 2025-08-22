import { ComponentProps } from "react";
import "./button.css";

interface Props extends ComponentProps<"button"> {}

export const Button: React.FC<Props> = ({ className, children }) => {
  return <button className={"button shadow " + (className ?? "")}>{children}</button>;
};
