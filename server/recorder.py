from datetime import datetime
from logging import Logger
from pathlib import Path
from audio_rw_thread import AudioReadWriteThread
from typing import Callable
from recording_state_notifier import RecordingStateNotifier

"""
Start and Stop recordings
"""
class Recorder:
    
    _rec_thread = None
    _data_dir = None

    def __init__(self, notifier: RecordingStateNotifier, data_dir: Path):
        self._notifier = notifier
        self._data_dir = data_dir

    def start(self) -> None:
        if self.get_is_recording():
            raise RuntimeError(
                "Can't start a recording while another one is running")
        file_path = self._generate_target_file_path()
        self._rec_thread = AudioReadWriteThread(file_path)
        self._rec_thread.start()
        self._notifier.notifyStarted(file_path)

    def stop(self) -> None:
        if not self.get_is_recording():
            raise RuntimeError("No recording is running")
        else:
            summary = self._rec_thread.stop()
            self._rec_thread = None
            self._notifier.notifyStopped(summary.file_name, summary.duration)


    def get_is_recording(self) -> bool:
        return self._rec_thread is not None and self._rec_thread.is_alive()

    def _generate_target_file_path(self) -> str:
        now = datetime.now()
        dir_name = now.strftime("%d-%m-%Y")
        file_name = now.strftime("%H-%M-%S-%f") + ".wav"
        target_path =  self._data_dir / dir_name / file_name
        if not target_path.parent.exists():
            target_path.parent.mkdir(parents = True)
        return target_path

    
