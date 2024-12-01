from dataclasses import dataclass
from typing import Callable, cast
from abc import ABC, abstractmethod
import types
from typing import Coroutine
import asyncio


@dataclass
class RecordingState:
    is_recording: bool
    file_name: str
    duration: int


class RecordingStateConsumerClass(ABC):
    @abstractmethod
    def on_state_change(self, state: RecordingState) -> None:
        pass


RecordingStateConsumer = RecordingStateConsumerClass | Coroutine[RecordingState, None, None]


class RecordingStateNotifier():

    _state_change_callbacks = []

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
        self._state_change_callbacks.append(callback)

    def _publish_state(self, is_recording: bool, file_name: str, duration: int):
        state = RecordingState(
            is_recording=is_recording,
            duration=duration,
            file_name=file_name
        )
        for cb in self._state_change_callbacks:
            try:
                if callable(cb):
                    cb(state)
                else:
                    loop = asyncio.get_event_loop()
                    consumer = cast(RecordingStateConsumer, cb)
                    loop.create_task(consumer.on_state_change(state))
            except BaseException as e:
                print(
                    f'[ERROR] Exception occured within state update callback: {str(e)}')
