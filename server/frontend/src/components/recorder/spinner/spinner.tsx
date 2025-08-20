import "./spinner.css";

interface Props {
  isRecording: boolean;
  className: string;
}

export const Spinner: React.FC<Props> = ({ isRecording, className }) => {
  return (
    <div
      className={
        "spinner " +
        (isRecording ? "active" : "") +
        (className ? " " + className : "")
      }
    ></div>
  );
};
