from datetime import datetime

from audio_rw_thread import AudioReadWriteThread

"""
Start and Stop recordings
"""
class Recorder:

    _rec_thread = None

    def start(self) -> None:
        if self.get_is_recording():
            raise RuntimeError(
                "Can't start a recording while another one is running")
        file_name = self._generate_recording_name()
        # TODO: Output directory?
        file_path = file_name
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

    def _generate_recording_name(self) -> str:
        return datetime.now().strftime("%d-%m-%Y-%H-%M-%S-%f") + ".wav"
