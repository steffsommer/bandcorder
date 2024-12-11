import threading
from queue import Queue
import sys
import soundfile as sf
import sounddevice as sd
from dataclasses import dataclass
from pathlib import Path
from typing import Callable
from recording_state_notifier import RecordingState
from datetime import datetime

NOTIFICATION_DELAY_SECONDS = 1


@dataclass
class RecordingSummary:
    file_name: str
    duration: int


"""
Relay microphone audio data to a file. Uses the microphone set as the
operating systems default microphone at object creation time.
"""


class AudioReadWriteThread(threading.Thread):

    _keep_recording = True
    mic_name: str
    mic_index: int
    sample_rate: int
    _state_update_callback: Callable[[RecordingState], None]
    _start_time: datetime
    _last_notification_time = datetime

    def __init__(
            self,
            out_file: Path,
            state_update_callback: Callable[[RecordingState], None]

    ):
        threading.Thread.__init__(self)
        if out_file is None:
            raise RuntimeError(
                "Can't create thread without valid output_file path set")
        self._out_File = out_file
        self.mic_index = sd.default.device[0]
        default_mic = sd.query_devices()[self.mic_index]
        self.mic_name = default_mic['name']
        self.sample_rate = int(default_mic['default_samplerate'])
        self._state_update_callback = state_update_callback

    def stop(self) -> RecordingSummary:
        self._keep_recording = False
        self.join()
        return RecordingSummary(file_name=self._out_File.name, duration=10)

    def run(self):
        self._start_time = datetime.now()
        self._last_notification_time = self._start_time
        q = Queue()
        with sf.SoundFile(self._out_File, mode='x', samplerate=self.sample_rate, channels=1, format="WAVEX", subtype=None) as file:
            callback = self._get_audio_write_callback(q)
            with sd.InputStream(samplerate=self.sample_rate, device=self.mic_index, channels=1, callback=callback):
                while self._keep_recording:
                    file.write(q.get())
                    self._notify()

    def _get_audio_write_callback(self, queue):
        def callback(indata, frames, time, status):
            """This is called (from a separate thread) for each audio block."""
            if status:
                print(status, file=sys.stderr)
            queue.put(indata.copy())
        return callback

    def _notify(self):
        now = datetime.now()
        time_delta = now - self._last_notification_time
        if time_delta.seconds < NOTIFICATION_DELAY_SECONDS:
            return
        duration = (now - self._start_time).seconds
        state = RecordingState(
            is_recording=True,
            duration=duration,
            file_name=self._out_File.name
        )
        self._state_update_callback(state)
        self._last_notification_time = now
