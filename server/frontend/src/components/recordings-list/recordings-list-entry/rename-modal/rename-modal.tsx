import { useEffect, useRef, useState } from "react";
import { FiFile, FiSave } from "react-icons/fi";
import { IoMdClose } from "react-icons/io";
import { RenameRecording } from "../../../../../wailsjs/go/services/FileSystemStorageService";
import { toastFailure, toastSuccess } from "../../../../services/toast-service/toast-service";
import { Button } from "../../../button/button";
import "./rename-modal.css";

interface Props {
  show?: boolean;
  onClose: () => void;
  initialFileValue: string;
  onRename: () => void;
}

export function RenameModal({ show, onClose, initialFileValue, onRename }: Props) {
  const dialogRef = useRef<HTMLDialogElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  const [isValid, setIsValid] = useState(true);

  useEffect(() => {
    if (show) {
      if (inputRef.current) {
        inputRef.current.value = (initialFileValue ?? "").split(".")[0];
      }
      dialogRef?.current?.showModal();
    } else {
      dialogRef?.current?.close();
    }
  }, [show]);

  const validateFileName = () => {
    const fileName = inputRef.current?.value ?? "";
    const regex = /^[a-zA-Z0-9_ -]+$/;
    const valid = regex.test(fileName);
    setIsValid(valid);
    return valid;
  };

  const handleSubmit = async () => {
    if (!isValid) return;
    try {
      const date = new Date().toISOString();
      const newName = (inputRef.current?.value ?? "") + ".wav";
      await RenameRecording(initialFileValue, newName, date);
      onRename();
      toastSuccess("Rename successful");
    } catch (e) {
      toastFailure("Rename failed");
    }
  };

  return (
    <dialog ref={dialogRef} className="rename-dialog" onClose={onClose}>
      <form method="dialog" onSubmit={handleSubmit}>
        <h2 className="rename-heading">
          <FiFile />
          <span>Rename file</span>
          <button type="button" className="close-btn" onClick={() => dialogRef?.current?.close()}>
            <IoMdClose className="rename-close-icon" />
          </button>
        </h2>
        <div className="rename-input-container">
          <label htmlFor="fileName">File name</label>
          <div style={{ display: "flex", alignItems: "center", gap: "0.25rem" }}>
            <input
              autoFocus
              ref={inputRef}
              id="file-name"
              name="fileName"
              onChange={validateFileName}
            />
            <span>.wav</span>
          </div>
          {!isValid && (
            <span style={{ color: "red", fontSize: "0.875rem" }} className="validation-message">
              File name not valid
            </span>
          )}
        </div>
        <Button className="save-btn" disabled={!isValid}>
          <FiSave />
          <span>Save</span>
        </Button>
      </form>
    </dialog>
  );
}
