from dataclasses import dataclass
from typing import Callable


@dataclass
class RecordingState:
    is_recording: bool
    file_name: str
    duration: int


class RecordingStateNotifier():

    _state_change_callbacks = []

    def notifyStarted(self, file_name: str) -> None:
        """Signal subscribers that a new recording was started"""
        self._publish_state(True, file_name, 0)

    def notifyStopped(self, file_name: str, duration: int) -> None:
        """Signal subscribers that the current recording was stopped"""
        self._publish_state(False, file_name, duration)

    def on_state_updates(self, callback: Callable[[RecordingState], None]):
        """Register a callback that gets executed at least every second with
        the current recording state.
        """
        self._state_change_callbacks.append(callback)

    def _publish_state(self, is_recording: bool, file_name: str, duration: int):
        state = RecordingState(
            is_recording=is_recording,
            duration=duration,
            file_name=file_name
        )
        for cb in self._state_change_callbacks:
            try:
                cb(state)
            except BaseException as e:
                print('[ERROR] Exception occured within state update callback')
