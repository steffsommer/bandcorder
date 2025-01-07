from dataclasses import dataclass
from typing import Callable, cast, Dict
from abc import ABC, abstractmethod
from interval_task import IntervalTask
from datetime import datetime
from logging import Logger

INTERVAL_SECONDS = 1


# Note: the duration property was chosen in favor of a start_time,
# to support future addition of Embedded devices that do not necessarily
# have synchronized time
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


class RecordingStateNotifier:

    _subscribers = []
    _state = RecordingState(is_recording=False, duration=0, file_name='-')
    _start_time: datetime = None

    def __init__(self, logger: Logger):
        self._logger = logger

    def notifyStarted(self, file_name: str) -> None:
        """Signal subscribers that a new recording was started"""
        self._state = RecordingState(
            is_recording=True,
            duration=0,
            file_name=file_name
        )
        self._start_time = datetime.now()
        self._publish()

    def notifyStopped(self) -> None:
        """Signal subscribers that the current recording was stopped"""
        duration = self._calculate_duration()
        self._state = RecordingState(
            is_recording=False,
            duration=duration,
            file_name='-'
        )
        self._start_time = None
        self._publish()

    def register_subscriber(self, callback: RecordingStateConsumer):
        """Register a callback that gets executed at least every second with
        the current recording state.
        """
        self._subscribers.append(callback)

    def start(self) -> None:
        task = IntervalTask(INTERVAL_SECONDS, self._publish_current_state)
        task.start()

    def _publish_current_state(self) -> None:
        self._state.duration = self._calculate_duration()
        self._publish()
        self._logger.debug(f'Published state to client {self._state}')

    def _calculate_duration(self):
        if self._start_time is None:
            return 0
        return (datetime.now() - self._start_time).seconds

    def _publish(self):
        for cb in self._subscribers:
            try:
                if callable(cb):
                    cb(self._state)
                else:
                    consumer = cast(RecordingStateConsumer, cb)
                    consumer.on_state_change(self._state)
            except BaseException as e:
                print(
                    f'[ERROR] Exception occured within state update callback: {str(e)}')
