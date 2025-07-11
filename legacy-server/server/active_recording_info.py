from dataclasses import dataclass
import tkinter as tk


@dataclass
class RecordingState():
    file_name: str
    seconds: int


class ActiveRecordingInfo(tk.Frame):

    def __init__(self, parent):
        super().__init__(parent)
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
        return tk.Label(self, font=("Arial", 20), text=icon)

    def update(self, state: RecordingState):
        self.file_label.config(text = state.file_name)
        self.time_label.config(text = state.seconds)
        

