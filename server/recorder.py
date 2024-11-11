import logging
from threading import Thread
from datetime import datetime
import queue
import sys
from audio_rw_thread import AudioReadWriteThread
import sounddevice as sd
import soundfile as sf


class Recorder:

    _rec_thread = None

    def start(self):
        if self.get_is_recording():
            raise RuntimeError(
                "Can't start a recording while another one is running")
        file_name = self._generate_recording_name()
        # TODO: Output directory?
        file_path = file_name
        self._rec_thread = AudioReadWriteThread(file_path)
        self._rec_thread.start()

    def stop(self):
        if not self.get_is_recording():
            raise RuntimeError("No recording is running")
        else:
            self._rec_thread.stop()

    def get_is_recording(self):
        return self._rec_thread is not None and self._rec_thread.is_alive()

    def _generate_recording_name(self):
        return datetime.now().strftime("%d-%m-%Y-%H-%M-%S-%f") + ".wav"
