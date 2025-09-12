import { useEffect, useRef } from "react";
import { FiSave, FiSettings } from "react-icons/fi";
import { IoMdClose } from "react-icons/io";
import "./settings-modal.css";
import { Button } from "../button/button";

interface Props {
  show?: boolean;
  onClose: () => void;
}

export function SettingsModal({ show, onClose }: Props) {
  const dialogRef = useRef<HTMLDialogElement>(null);

  useEffect(() => {
    if (show) {
      dialogRef?.current?.showModal();
    } else {
      dialogRef?.current?.close();
    }
  }, [show]);

  const handleSubmit = () => {
    console.log("Submit not implemented yet");
  };

  return (
    <dialog ref={dialogRef} className="settings-dialog" onClose={onClose}>
      <form method="dialog" onSubmit={handleSubmit}>
        <h2 className="settings-heading">
          <FiSettings />
          <span>Settings</span>
          <button
            type="button"
            className="close-btn"
            onClick={() => dialogRef?.current?.close()}
          >
            <IoMdClose className="settings-close-icon" />
          </button>
        </h2>
        <div className="settings-list">
          <label htmlFor="recording-directory">Recordings directory</label>
          <input id="recording-directory" />
        </div>
        <Button className="save-btn">
          <FiSave />
          <span>Save</span>
        </Button>
      </form>
    </dialog>
  );
}
