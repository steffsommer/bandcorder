import tkinter as tk
import customtkinter as ctk
from tkinter import ttk
from file_storage_service import FileStorageService
from recording_state_notifier import RecordingStateNotifier, RecordingState
import time
import datetime


class RecordingsTreeView(ctk.CTkFrame):

    _last_state: RecordingState = None

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
            text="Todays recordings",
            font=("Arial", 40),
        )

        # placement
        self._treeview.grid(row=1, column=0)
        self._recordings_list_label.grid(row=0, column=0)

        notifier.register_subscriber(self._set_recordings)
        self._update_list()

    def _set_recordings(self, recording_state: RecordingState) -> None:
        if self._last_state is not None and self._last_state.is_recording and not recording_state.is_recording:
            self._update_list()
        self._last_state = recording_state

    def _get_configured_treeview(self) -> ttk.Treeview:
        treeview = ttk.Treeview(self, columns=("duration"), height=20)
        treeview.heading("#0", text="File")
        treeview.column("#0", width=300)
        treeview.heading("#1", text="Duration")
        treeview.column("#1", width=150, anchor=tk.CENTER)
        return treeview

    def _update_list(self):
        recordings = self._storage_service.get_todays_recordings()
        for item in self._treeview.get_children():
            self._treeview.delete(item)
        recordings.reverse()
        for recording in recordings:
            duration_str = str(datetime.timedelta(
                seconds=int(recording.duration)))
            self._treeview.insert(
                "", tk.END, text=f'📼{recording.name}', values=(duration_str))
