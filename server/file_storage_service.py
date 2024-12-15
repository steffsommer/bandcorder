from pathlib import Path
from datetime import datetime
import os
from os.path import isfile, join
from dataclasses import dataclass
import soundfile as sf


@dataclass
class Recording:
    name: str
    duration: float

# Obtain writable file handles for files based on current time
# Notify subscribers when new recordings have been started


class FileStorageService:

    def __init__(self, data_dir: Path):
        self._data_dir = data_dir

    def create_writable_wav_file(self) -> Path:
        path = self._create_file_internal()
        return path

    def _create_file_internal(self) -> Path:
        now = datetime.now()
        dir_name = now.strftime("%d-%m-%Y")
        file_name = now.strftime("%H-%M-%S-%f") + ".wav"
        target_path = self._data_dir / dir_name / file_name
        if not target_path.parent.exists():
            target_path.parent.mkdir(parents=True)
        return target_path

    def get_todays_recordings(self) -> list[Recording]:
        now = datetime.now()
        dir_name = now.strftime("%d-%m-%Y")
        todays_dir = self._data_dir.joinpath(dir_name)
        if not todays_dir.exists():
            return []

        recording_files = [os.path.join(todays_dir, file) for file in os.listdir(todays_dir)]
        soundfiles = [sf.SoundFile(file) for file in recording_files]
        recordings = []
        for soundfile in soundfiles:
            duration = soundfile.frames / soundfile.samplerate
            filename = os.path.basename(soundfile.name)
            recording = Recording(filename, duration)
            recordings.append(recording)
        return recordings
