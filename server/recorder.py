from datetime import datetime
from logging import Logger
from pathlib import Path
from audio_rw_thread import AudioReadWriteThread
from dataclasses import dataclass
from typing import Callable

@dataclass
class RecordingState:
    is_recording: bool
    file_name: str
    duration: int


"""
Start and Stop recordings
"""
class Recorder:
    
    _rec_thread = None
    _data_dir = None
    _state_change_callbacks = []

    def __init__(self, data_dir: Path):
        self._data_dir = data_dir

    def start(self) -> None:
        if self.get_is_recording():
            raise RuntimeError(
                "Can't start a recording while another one is running")
        file_path = self._generate_target_file_path()
        self._rec_thread = AudioReadWriteThread(file_path)
        self._rec_thread.start()
        state_update = RecordingState(
            is_recording=True,
            duration=0,
            file_name=file_path
        )
        self._publish_state_update(state_update)

    def stop(self) -> None:
        if not self.get_is_recording():
            raise RuntimeError("No recording is running")
        else:
            summary = self._rec_thread.stop()
            self._rec_thread = None
            state_update = RecordingState(
                is_recording=False,
                duration=summary.duration,
                file_name=summary.file_name
            )
            self._publish_state_update(state_update)

    def _publish_state_update(self, state: RecordingState):
        for cb in self._state_change_callbacks:
            try:
                cb(state)
            except BaseException as e:
                print('[ERROR] Exception occured within state update callback')

    def get_is_recording(self) -> bool:
        return self._rec_thread is not None and self._rec_thread.is_alive()

    def on_state_updates(self, callback: Callable[[RecordingState], None]):
        self._state_change_callbacks.append(callback)

    def _generate_target_file_path(self) -> str:
        now = datetime.now()
        dir_name = now.strftime("%d-%m-%Y")
        file_name = now.strftime("%H-%M-%S-%f") + ".wav"
        target_path =  self._data_dir / dir_name / file_name
        if not target_path.parent.exists():
            target_path.parent.mkdir(parents = True)
        return target_path

    
