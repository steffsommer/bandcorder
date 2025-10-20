import { useEffect, useRef } from "react";
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
}

export function RenameModal({ show, onClose, initialFileValue }: Props) {
  const dialogRef = useRef<HTMLDialogElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (show) {
      if (inputRef.current) {
        inputRef.current.value = initialFileValue;
      }
      dialogRef?.current?.showModal();
    } else {
      dialogRef?.current?.close();
    }
  }, [show]);

  const handleSubmit = async () => {
    try {
      const date = new Date().toISOString();
      RenameRecording(initialFileValue, inputRef.current?.value ?? "", date);
      toastSuccess("Rename successful");
    } catch (e) { }
    toastFailure("Rename failed");
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
        <div className="rename-list">
          <label htmlFor="fileName">File name</label>
          <input autoFocus ref={inputRef} id="file-name" name="fileName" />
        </div>
        <Button className="save-btn">
          <FiSave />
          <span>Save</span>
        </Button>
      </form>
    </dialog>
  );
}
