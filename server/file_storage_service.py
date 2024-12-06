from pathlib import Path
from typing import Callable
from datetime import datetime
from os import listdir
from os.path import isfile, join
from logging import Logger
from dataclasses import dataclass


@dataclass
class Recording:
    name: str

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
        return [Recording(file) for file in listdir(todays_dir)
                if isfile(join(todays_dir, file))]
