from dataclasses import dataclass
import tkinter as tk
from recording_state_notifier import RecordingStateNotifier, RecordingState


class ActiveRecordingInfo(tk.Frame):

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
        self.file_label.grid(row=0, column=1)
        self.time_label.grid(row=1, column=1)

    def _get_label(self, icon: str) -> tk.Label:
        return tk.Label(self, font=("Arial", 20), text=icon, justify='left')

    def update(self, state: RecordingState):
        if state.is_recording:
            self.grid()
        else:
            self.grid_remove()
        self.file_label.config(text = state.file_name)
        self.time_label.config(text = f'{state.duration}s')
        