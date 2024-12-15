import tkinter as tk
import customtkinter as ctk
from tkinter import ttk
from file_storage_service import FileStorageService
from recording_state_notifier import RecordingStateNotifier, RecordingState


class RecordingsTreeView(ctk.CTkFrame):

    def __init__(
            self,
            parent,
            notifier: RecordingStateNotifier,
            storage_service: FileStorageService
    ):
        super().__init__(parent)
        self._storage_service = storage_service

        # component definitions
        self._treeview = self._get_configured_treeview()
        self._recordings_list_label = ctk.CTkLabel(
            self,
            text="Todays recording",
            font=("Arial", 40),
        )

        # placement
        self._treeview.grid(row=1, column=0)
        self._recordings_list_label.grid(row=0, column=0)

        notifier.register_subscriber(self._set_recordings)
        self._update_list()

    def _set_recordings(self, recording_state: RecordingState) -> None:
        if recording_state.is_recording:
            return
        self._update_list()

    def _get_configured_treeview(self) -> ttk.Treeview:
        treeview = ttk.Treeview(self, columns=("size", "lastmod"), height=20)
        treeview.heading("#0", text="File")
        return treeview

    def _update_list(self):
        recordings = self._storage_service.get_todays_recordings()
        for item in self._treeview.get_children():
            self._treeview.delete(item)
        for recording in recordings:
            self._treeview.insert("", tk.END, text=recording.name)