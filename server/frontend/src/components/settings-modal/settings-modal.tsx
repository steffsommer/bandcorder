import { useEffect, useRef } from "react";
import { FiSettings } from "react-icons/fi";
import { FaRegWindowClose } from "react-icons/fa";
import "./settings-modal.css";

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

  return (
    <dialog ref={dialogRef} className="settings-dialog" onClose={onClose}>
      <form method="dialog">
        <h2 className="settings-heading">
          <FiSettings />
          <span>Settings</span>
          <button type="submit" className="close-btn">
            <FaRegWindowClose className="settings-close-icon" />
          </button>
        </h2>
      </form>
    </dialog>
  );
}
