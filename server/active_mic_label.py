import tkinter as tk
from audio_rw_thread import AudioReadWriteThread
from pathlib import Path

NO_MIC_TEXT = 'NO ACTIVE MIC'
NO_MIC_BACKGROUND_COLOR = 'red'
DEFAULT_BACKGROUND_COLOR = 'white'
UPDATE_PERIOD_MS = 1000


class ActiveMicLabel(tk.Label):

    def __init__(self, parent: tk.Widget):
        super().__init__(
            parent,
            font=("Arial", 20),
            pady=20,
        )
        self.update_active_mic()

    def update_active_mic(self) -> None:
        mic_name = None
        try:
            rw_thread = AudioReadWriteThread(Path())
            mic_name = rw_thread.mic_name
        except Exception as e:
            print(e)
        if mic_name is None or mic_name == '':
            self.config(
                text=NO_MIC_TEXT,
                background=NO_MIC_BACKGROUND_COLOR,
            )
            return
        self.config(
            text=mic_name,
            background=DEFAULT_BACKGROUND_COLOR,
        )
        self.after(UPDATE_PERIOD_MS, self.update_active_mic)
