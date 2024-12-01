import tkinter as tk

NO_MIC_TEXT = 'NO ACTIVE MIC'
NO_MIC_BACKGROUND_COLOR = 'red'

DEFAULT_BACKGROUND_COLOR = 'white'


class ActiveMicLabel(tk.Label):

    def __init__(self, parent: tk.Widget, microphone: str):
        super().__init__(
            parent,
            font=("Arial", 20),
            pady=20,
        )
        self.set_active_mic(microphone)

    def set_active_mic(self, microphone: str) -> None:
        if microphone is None or microphone == '':
            self.config(
                text=NO_MIC_TEXT,
                background=NO_MIC_BACKGROUND_COLOR,
            )
            return
        self.config(
            text=microphone,
            background=DEFAULT_BACKGROUND_COLOR,
        )
