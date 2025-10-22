import { Card } from "../../card/card";
import "./recordings-list-entry.css";
import { FaMusic } from "react-icons/fa";
import { MdOutlineAccessTimeFilled } from "react-icons/md";
import { FaEdit, FaTrash } from "react-icons/fa";
import { Button } from "../../button/button";
import { models } from "../../../../wailsjs/go/models";
import { TimeUtils } from "../../../utils/time";
import { RenameModal } from "./rename-modal/rename-modal";
import { useState } from "react";

interface Props {
  recording: models.RecordingInfo;
  onRename: () => void;
}

export const RecordingsListEntry: React.FC<Props> = ({ recording, onRename }) => {
  const [showRenameModal, setShowRenameModal] = useState(false);
  return (
    <Card className="entry-card">
      <div className="row">
        <div className="note-icon-container">
          <FaMusic size="2em" />
        </div>
        <h3 className="recording-title">{recording.fileName}</h3>
        <span className="duration">
          <MdOutlineAccessTimeFilled size="2em" />
          <span className="duration-str">{TimeUtils.toMMSS(recording.durationSeconds)}</span>
        </span>
        <Button className="list-btn edit-btn" onClick={() => setShowRenameModal(true)}>
          <FaEdit />
        </Button>
        <Button className="list-btn">
          <FaTrash />
        </Button>
      </div>
      <RenameModal
        show={showRenameModal}
        initialFileValue={recording.fileName}
        onClose={() => setShowRenameModal(false)}
        onRename={onRename}
      />
    </Card>
  );
};
