import { Card } from "../../card/card";
import "./recordings-list-entry.css";
import { FaMusic } from "react-icons/fa";
import { MdOutlineAccessTimeFilled } from "react-icons/md";
import { FaEdit, FaTrash } from "react-icons/fa";
import { Button } from "../../button/button";
import { models } from "../../../../wailsjs/go/models";
import { TimeUtils } from "../../../utils/time";

interface Props {
  recording: models.RecordingInfo;
}

export const RecordingsListEntry: React.FC<Props> = ({ recording }) => {
  return (
    <Card className="entry-card">
      <div className="row">
        <div className="note-icon-container">
          <FaMusic size="2em" />
        </div>
        <h3 className="recording-title">{recording.FileName}</h3>
        <span className="duration">
          <MdOutlineAccessTimeFilled size="2em" />
          <span className="duration-str">{TimeUtils.toMinutesSecondsStr(recording.DurationSeconds)}</span>
        </span>
        <Button className="list-btn edit-btn">
          <FaEdit />
        </Button>
        <Button className="list-btn">
          <FaTrash />
        </Button>
      </div>
    </Card>
  );
};
