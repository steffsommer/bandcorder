import threading
from queue import Queue
import sys
import soundfile as sf
import sounddevice as sd

"""
Relay microphone audio data to a file
"""
class AudioReadWriteThread(threading.Thread):

    _keep_recording = True

    def __init__(self, out_file):
        threading.Thread.__init__(self)
        if out_file is None:
            raise RuntimeError(
                "Can't create thread without valid output_file path set")
        self.out_file = out_file

    def stop(self):
        self._keep_recording = False
        self.join()

    def run(self):
        q = Queue()
        mic_index = sd.default.device[0]
        default_mic = sd.query_devices()[mic_index]
        sample_rate = default_mic['default_samplerate']
        sample_rate_int = int(sample_rate)
        with sf.SoundFile(self.out_file, mode='x', samplerate=sample_rate_int, channels=1, format="WAVEX", subtype=None) as file:
            callback = self._get_callback(q)
            with sd.InputStream(samplerate=sample_rate, device=mic_index, channels=1, callback=callback):
                while self._keep_recording:
                    file.write(q.get())

    def _get_callback(self, queue):
        def callback(indata, frames, time, status):
            """This is called (from a separate thread) for each audio block."""
            if status:
                print(status, file=sys.stderr)
            queue.put(indata.copy())
        return callback
