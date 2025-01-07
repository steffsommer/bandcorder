from dataclasses import dataclass
import tkinter as tk
import customtkinter as ctk
from recording_state_notifier import RecordingStateNotifier, RecordingState
import datetime


class ActiveRecordingInfo(ctk.CTkFrame):

    def __init__(
            self,
            parent: tk.Widget,
            notifier: RecordingStateNotifier,
    ):
        super().__init__(parent)
        self._setup_layout()
        notifier.register_subscriber(self.update)
        self.grid_remove()

    def _setup_layout(self):
        # component definitions
        self.file_icon_label = self._get_label('📼')
        self.time_icon_label = self._get_label('⏰️')
        self.file_label = self._get_label('')
        self.time_label = self._get_label('')
        # component placement
        self.file_icon_label.grid(row=0, column=0)
        self.time_icon_label.grid(row=1, column=0)
        self.file_label.grid(row=0, column=1, sticky='W')
        self.time_label.grid(row=1, column=1, sticky='W')

    def _get_label(self, icon: str) -> ctk.CTkLabel:
        return ctk.CTkLabel(self, font=("Arial", 30), text=icon, justify='left', padx=20, pady=20)

    def update(self, state: RecordingState):
        if state.is_recording:
            self.grid()
        else:
            self.grid_remove()
        self.file_label.configure(text=state.file_name)
        duration_str = str(datetime.timedelta(
            seconds=int(state.duration)))
        self.time_label.configure(text=duration_str)
