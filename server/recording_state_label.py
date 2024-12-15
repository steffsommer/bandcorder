import tkinter as tk
import customtkinter as ctk
from recording_state_notifier import RecordingStateNotifier, RecordingState

ON_AIR_BACKGROUND_COLOR = 'green'
ON_AIR_TEXT = 'ON AIR'
ON_AIR_TEXT_COLOR = 'black'

IDLE_BACKGROUND_COLOR = 'red'
IDLE_TEXT = 'IDLE'
IDLE_TEXT_COLOR = 'white'


class RecordingStateLabel(ctk.CTkLabel):
    def __init__(self, parent, notifier: RecordingStateNotifier):
        super().__init__(
            parent,
            font=("Arial", 100),
            width=800
        )
        self._set_idle()
        notifier.register_subscriber(self._process_status_updates)
    
    def _process_status_updates(self, state: RecordingState):
        if state.is_recording:
            self._set_on_air()
        else:
            self._set_idle()

    def _set_idle(self) -> None:
        self.configure(
            text=IDLE_TEXT,
            bg_color=IDLE_BACKGROUND_COLOR,
            text_color=IDLE_TEXT_COLOR
        )

    def _set_on_air(self) -> None:
        self.configure(
            text=ON_AIR_TEXT,
            bg_color=ON_AIR_BACKGROUND_COLOR,
            text_color=ON_AIR_TEXT_COLOR
        )
