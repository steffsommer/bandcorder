import logging
from datetime import datetime
import queue
import sys
import sounddevice as sd
import soundfile as sf


class Recorder:

    _isRecording = False
    _logger = logging.getLogger("Recorder")

    def start(self):
        if self._isRecording:
            raise RuntimeError(
                "Can't start a recording while another one is running")
        file_name = self._generate_recording_name()
        # TODO: Output directory?
        file_path = file_name
        self._logger.info("Starting recording {}", file_path)
        self._isRecording = True
        self._record_to_file(file_path)

    def stop(self):
        pass

    def get_state(self):
        pass

    def _generate_recording_name(self):
        return datetime.now().strftime("%d-%m-%Y-%H-%M-%S-%f")

    def _get_callback(self, queue):
        def callback(indata, frames, time, status):
            """This is called (from a separate thread) for each audio block."""
            if status:
                print(status, file=sys.stderr)
            queue.put(indata.copy())
        return callback

    def _record_to_file(self, file_path):
        q = queue.Queue()
        try:
            mic_index = sd.default.device[0]
            default_mic = sd.query_devices()[mic_index]
            sample_rate = default_mic['default_samplerate']
            sample_rate_int = int(sample_rate)
            # Make sure the file is opened before recording anything:
            with sf.SoundFile(file_path, mode='x', samplerate=sample_rate_int, channels=1, format="WAVEX", subtype=None) as file:
                callback = self._get_callback(q)
                with sd.InputStream(samplerate=sample_rate, device=mic_index, channels=1, callback=callback):
                    while True:
                        print("received audio data")
                        # file.write(q.get())
        except KeyboardInterrupt:
            print('\nRecording finished: ' + repr(file_path))
