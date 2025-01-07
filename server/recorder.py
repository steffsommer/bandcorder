from audio_rw_thread import AudioReadWriteThread
from recording_state_notifier import RecordingStateNotifier
from file_storage_service import FileStorageService

"""
Start and Stop recordings
"""


class Recorder:

    _rec_thread = None

    def __init__(
            self,
            notifier: RecordingStateNotifier,
            storage_service: FileStorageService,
    ):
        self._notifier = notifier
        self._storage_service = storage_service

    def start(self) -> None:
        if self.is_recording():
            raise RuntimeError(
                "Can't start a recording while another one is running")
        file_path = self._storage_service.create_writable_wav_file()
        self._rec_thread = AudioReadWriteThread(file_path)
        self._rec_thread.start()
        self._notifier.notifyStarted(file_path.name)

    def stop(self) -> None:
        if not self.is_recording():
            raise RuntimeError("No recording is running")
        else:
            summary = self._rec_thread.stop()
            self._rec_thread = None
            self._notifier.notifyStopped()

    def is_recording(self) -> bool:
        return self._rec_thread is not None and self._rec_thread.is_alive()
