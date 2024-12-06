import threading
from queue import Queue
import sys
import soundfile as sf
import sounddevice as sd
from dataclasses import dataclass
from pathlib import Path


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

    def __init__(self, out_file: Path):
        threading.Thread.__init__(self)
        if out_file is None:
            raise RuntimeError(
                "Can't create thread without valid output_file path set")
        self._out_File = out_file
        self.mic_index = sd.default.device[0]
        default_mic = sd.query_devices()[self.mic_index]
        self.mic_name = default_mic['name']
        self.sample_rate = int(default_mic['default_samplerate'])

    def stop(self) -> RecordingSummary:
        self._keep_recording = False
        self.join()
        return RecordingSummary(file_name=self._out_File.name, duration=10)

    def run(self):
        q = Queue()
        with sf.SoundFile(self._out_File, mode='x', samplerate=self.sample_rate, channels=1, format="WAVEX", subtype=None) as file:
            callback = self._get_callback(q)
            with sd.InputStream(samplerate=self.sample_rate, device=self.mic_index, channels=1, callback=callback):
                while self._keep_recording:
                    file.write(q.get())

    def _get_callback(self, queue):
        def callback(indata, frames, time, status):
            """This is called (from a separate thread) for each audio block."""
            if status:
                print(status, file=sys.stderr)
            queue.put(indata.copy())
        return callback
