import "./spinner.css";

interface Props {
  isRecording: boolean;
}


export const Spinner: React.FC<Props> = ({isRecording}) => {
  return (
    <div className={"spinner " + (isRecording ? "active" : "")}></div>
  );
};
