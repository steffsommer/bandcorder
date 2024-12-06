import tkinter as tk
import threading
from audio_rw_thread import AudioReadWriteThread
import time
from pathlib import Path

NO_MIC_TEXT = 'NO ACTIVE MIC'
NO_MIC_BACKGROUND_COLOR = 'red'
DEFAULT_BACKGROUND_COLOR = 'white'


class ActiveMicLabel(tk.Label):

    def __init__(self, parent: tk.Widget):
        super().__init__(
            parent,
            font=("Arial", 20),
            pady=20,
        )
        update_thread = threading.Thread(target=self.update_active_mic)
        update_thread.start()

    def update_active_mic(self) -> None:
        while True:
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
            time.sleep(1)
