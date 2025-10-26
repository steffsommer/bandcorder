import { useEffect, useRef } from "react";
import { FiCheck, FiMinusCircle } from "react-icons/fi";
import { IoMdClose, IoMdWarning } from "react-icons/io";
import { Button } from "../button/button";
import "./confirmation-modal.css";

interface Props {
  show?: boolean;
  onClose: () => void;
  text: string;
  onAccept: () => void;
}

export function ConfirmationModal({ show, onClose, text, onAccept }: Props) {
  const dialogRef = useRef<HTMLDialogElement>(null);

  useEffect(() => {
    if (show) {
      dialogRef?.current?.showModal();
    } else {
      dialogRef?.current?.close();
    }
  }, [show]);

  return (
    <dialog ref={dialogRef} className="confirmation-dialog" onClose={onClose}>
      <h2 className="confirmation-heading">
        <IoMdWarning />
        <span>Confirm</span>
        <button type="button" className="close-btn" onClick={() => dialogRef?.current?.close()}>
          <IoMdClose className="confirmation-close-icon" />
        </button>
      </h2>
      <span className="confirmation-text">{text}</span>
      <div className="confirm-actions">
        <Button className="action-btn confirm-btn" onClick={onAccept}>
          <FiCheck />
          <span>Yes</span>
        </Button>
        <Button className="action-btn reject-btn" onClick={onClose}>
          <FiMinusCircle />
          <span>No</span>
        </Button>
      </div>
    </dialog>
  );
}
