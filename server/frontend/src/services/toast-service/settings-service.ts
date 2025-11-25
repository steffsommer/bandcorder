import { models } from "../../../wailsjs/go/models";
import { Load, Save } from "../../../wailsjs/go/services/SettingsService";
import { toastFailure, toastSuccess } from "./toast-service";

export async function loadSettings(): Promise<models.Settings> {
  try {
    const settings = await Load();
    return settings;
  } catch (e) {
    toastFailure("Failed to load settings");
    throw e;
  }
}

export async function saveSettings(settings: models.Settings): Promise<void> {
  try {
    await Save(settings);
    toastSuccess("Settings updated successfully");
  } catch (e) {
    toastFailure("Failed to save settings");
    throw e;
  }
}
