from dataclasses import dataclass
from typing import Callable, cast, Dict
from abc import ABC, abstractmethod
import asyncio

INTERVAL_SECONDS = 1

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
    _state = RecordingState(is_recording=False, duration=0, file_name='')

    def notifyStarted(self, file_name: str) -> None:
        """Signal subscribers that a new recording was started"""
        state = RecordingState(
            is_recording=True,
            duration=0,
            file_name=file_name
        )
        self.publish(state)

    def notifyStopped(self, file_name: str, duration: int) -> None:
        """Signal subscribers that the current recording was stopped"""
        state = RecordingState(
            is_recording=False,
            duration=duration,
            file_name=file_name
        )
        self.publish(state)

    def register_subscriber(self, callback: RecordingStateConsumer):
        """Register a callback that gets executed at least every second with
        the current recording state.
        """
        self._subscribers.append(callback)
    
    def start(self) -> None:
        loop = asyncio.get_event_loop()
        self.loop_send_task = loop.create_task(self._send_periodically())

    async def _send_periodically(self) -> None:
        while True:
            self.on_state_change(self._state)
            self._logger.debug(f'Published state to client {self._state}')
            await asyncio.sleep(INTERVAL_SECONDS)


    def publish(self, event: RecordingState):
        for cb in self._subscribers:
            try:
                if callable(cb):
                    cb(event)
                else:
                    consumer = cast(RecordingStateConsumer, cb)
                    consumer.on_state_change(event)
            except BaseException as e:
                print(
                    f'[ERROR] Exception occured within state update callback: {str(e)}')
