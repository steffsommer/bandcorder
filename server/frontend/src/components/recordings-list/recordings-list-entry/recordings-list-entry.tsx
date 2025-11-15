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
import { ConfirmationModal } from "../../confirmation-modal/confirmation-modal";
import { DeleteRecording } from "../../../../wailsjs/go/facades/FileSystemStorageFacade";
import { toastFailure, toastSuccess } from "../../../services/toast-service/toast-service";

interface Props {
  recording: models.RecordingInfo;
  onChange: () => void;
}

export const RecordingsListEntry: React.FC<Props> = ({ recording, onChange }) => {
  const [showRenameModal, setShowRenameModal] = useState(false);
  const [showDeleteModal, setShowDeleteModal] = useState(false);

  async function deleteRecording() {
    const date = new Date().toISOString();
    try {
      await DeleteRecording(recording.fileName, date);
      toastSuccess("Recording deleted");
    } catch (e) {
      toastFailure("Failed to delete recording");
    }
    onChange();
  }

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
        <Button className="list-btn" onClick={() => setShowDeleteModal(true)}>
          <FaTrash />
        </Button>
      </div>
      <RenameModal
        show={showRenameModal}
        initialFileValue={recording.fileName}
        onClose={() => setShowRenameModal(false)}
        onRename={onChange}
      />
      <ConfirmationModal
        show={showDeleteModal}
        onAccept={deleteRecording}
        text="Do you really want to delete this recording?"
        onClose={() => setShowDeleteModal(false)}
      />
    </Card>
  );
};
