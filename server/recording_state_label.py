import tkinter as tk

ON_AIR_BACKGROUND_COLOR = 'green'
ON_AIR_TEXT = 'ON AIR 🎤'
ON_AIR_TEXT_COLOR = 'black'

IDLE_BACKGROUND_COLOR = 'red'
IDLE_TEXT = 'IDLE'
IDLE_TEXT_COLOR = 'white'


class RecordingStateLabel(tk.Label):
    def __init__(self, parent):
        super().__init__(
            parent,
            font=("Arial", 100),
            padx=100,
            pady=100,
            width=8
        )
        self.set_idle()

    def set_idle(self) -> None:
        self.config(
            text=IDLE_TEXT,
            background=IDLE_BACKGROUND_COLOR,
            fg=IDLE_TEXT_COLOR
        )

    def set_on_air(self) -> None:
        self.config(
            text=ON_AIR_TEXT,
            background=ON_AIR_BACKGROUND_COLOR,
            fg=ON_AIR_TEXT_COLOR
        )
