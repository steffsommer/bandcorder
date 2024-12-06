from dataclasses import dataclass
from typing import Callable, cast, Dict
from abc import ABC, abstractmethod


@dataclass
class RecordingState:
    is_recording: bool
    file_name: str
    duration: int

    def toSerializableDict(self) -> Dict:
        return {
            'isRecording': self.is_recording,
            'fileName': self.file_name,
            'duration': self.duration,
        }


class RecordingStateConsumerClass(ABC):
    @abstractmethod
    def on_state_change(self, state: RecordingState) -> None:
        pass


RecordingStateConsumer = RecordingStateConsumerClass | Callable[[
    RecordingState], None]


class RecordingStateNotifier():

    _subscribers = []

    def notifyStarted(self, file_name: str) -> None:
        """Signal subscribers that a new recording was started"""
        self._publish_state(True, file_name, 0)

    def notifyStopped(self, file_name: str, duration: int) -> None:
        """Signal subscribers that the current recording was stopped"""
        self._publish_state(False, file_name, duration)

    def register_subscriber(self, callback: RecordingStateConsumer):
        """Register a callback that gets executed at least every second with
        the current recording state.
        """
        self._subscribers.append(callback)

    def _publish_state(self, is_recording: bool, file_name: str, duration: int):
        state = RecordingState(
            is_recording=is_recording,
            duration=duration,
            file_name=file_name
        )
        for cb in self._subscribers:
            try:
                if callable(cb):
                    cb(state)
                else:
                    consumer = cast(RecordingStateConsumer, cb)
                    consumer.on_state_change(state)
            except BaseException as e:
                print(
                    f'[ERROR] Exception occured within state update callback: {str(e)}')
