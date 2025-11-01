import { ReactNode } from "react";
import "./button.css";

interface Props {
  className?: string;
  disabled?: boolean;
  children?: ReactNode;
  onClick?: () => any;
}

export const Button: React.FC<Props> = ({ className, children, onClick, disabled }) => {
  return (
    <button disabled={disabled} onClick={onClick} className={"button shadow " + (className ?? "")}>
      {children}
    </button>
  );
};
