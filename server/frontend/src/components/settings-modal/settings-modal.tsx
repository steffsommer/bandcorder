import { useEffect, useRef, useState } from "react";
import { FiSave, FiSettings } from "react-icons/fi";
import { IoMdClose } from "react-icons/io";
import "./settings-modal.css";
import { Button } from "../button/button";
import { loadSettings, saveSettings } from "../../services/toast-service/settings-service";
import { models } from "../../../wailsjs/go/models";

interface Props {
  show?: boolean;
  onClose: () => void;
}

export function SettingsModal({ show, onClose }: Props) {
  const dialogRef = useRef<HTMLDialogElement>(null);
  const [settings, setSettings] = useState<models.Settings>({
    RecordingsDirectory: "",
  });

  useEffect(() => {
    if (show) {
      const loadAndShow = async () => {
        const settings = await loadSettings();
        setSettings(settings);
        dialogRef?.current?.showModal();
      };
      loadAndShow();
    } else {
      dialogRef?.current?.close();
    }
  }, [show]);

  const updateValue = (name: keyof models.Settings, value: string) => {
    const updated: models.Settings = {
      ...settings,
      [name]: value,
    };
    setSettings(updated);
  };

  const handleSubmit = async () => {
    await saveSettings(settings);
  };

  return (
    <dialog ref={dialogRef} className="settings-dialog" onClose={onClose}>
      <form method="dialog" onSubmit={handleSubmit}>
        <h2 className="settings-heading">
          <FiSettings />
          <span>Settings</span>
          <button type="button" className="close-btn" onClick={() => dialogRef?.current?.close()}>
            <IoMdClose className="settings-close-icon" />
          </button>
        </h2>
        <div className="settings-list">
          <label htmlFor="recording-directory">Recordings directory</label>
          <input
            id="recording-directory"
            name="recording-directory"
            value={settings?.RecordingsDirectory}
            onChange={(e) => updateValue("RecordingsDirectory", e.target.value)}
          />
        </div>
        <Button className="save-btn">
          <FiSave />
          <span>Save</span>
        </Button>
      </form>
    </dialog>
  );
}
