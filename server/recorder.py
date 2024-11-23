from datetime import datetime
from logging import Logger
from pathlib import Path

from audio_rw_thread import AudioReadWriteThread

"""
Start and Stop recordings
"""
class Recorder:
    
    _rec_thread = None
    _data_dir = None

    def __init__(self, data_dir: Path):
        self._data_dir = data_dir

    def start(self) -> None:
        if self.get_is_recording():
            raise RuntimeError(
                "Can't start a recording while another one is running")
        file_path = self._generate_target_file_path()
        self._rec_thread = AudioReadWriteThread(file_path)
        self._rec_thread.start()

    def stop(self) -> None:
        if not self.get_is_recording():
            raise RuntimeError("No recording is running")
        else:
            self._rec_thread.stop()
            self._rec_thread = None

    def get_is_recording(self) -> bool:
        return self._rec_thread is not None and self._rec_thread.is_alive()

    def _generate_target_file_path(self) -> str:
        file_name = datetime.now().strftime("%d-%m-%Y-%H-%M-%S-%f") + ".wav"
        return self._data_dir / file_name

    
