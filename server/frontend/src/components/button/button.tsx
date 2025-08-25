import { ReactNode } from "react";
import "./button.css";

interface Props {
  className?: string;
  children?: ReactNode;
  onClick?: () => any;
}

export const Button: React.FC<Props> = ({ className, children, onClick }) => {
  return (
    <button onClick={onClick} className={"button shadow " + (className ?? "")}>{children}</button>
  );
};
