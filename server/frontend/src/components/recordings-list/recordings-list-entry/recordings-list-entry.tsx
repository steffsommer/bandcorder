import { Card } from "../../card/card";
import { Btn } from "../../icon-btn/icon-btn";
import "./recordings-list-entry.css";
import { FaMusic } from "react-icons/fa";
import { MdOutlineAccessTimeFilled } from "react-icons/md";
import { FaEdit, FaTrash } from "react-icons/fa";
import { RecorderButton } from "../../recorder/recorder-button/recorder-button";

export interface Recording {
  name: string;
  duration: string;
}

interface Props {
  recording: Recording;
}

export const RecordingsListEntry: React.FC<Props> = ({ recording }) => {
  return (
    <Card className="entry-card">
      <div className="row">
        <div className="note-icon-container">
          <FaMusic size="2em" />
        </div>
        <h3 className="recording-title">{recording.name}</h3>
        <span className="duration">
          <MdOutlineAccessTimeFilled size="2em" />
          <span>{recording.duration}</span>
        </span>
        <RecorderButton bg="red" className="list-btn edit-btn">
          <FaEdit />
        </RecorderButton>
        <RecorderButton className="list-btn">
          <FaTrash />
        </RecorderButton>
      </div>
    </Card>
  );
};
