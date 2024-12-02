import socketio
from recorder import Recorder
from logging import Logger

NAMESPACE = '/'

class RecordingController(socketio.AsyncNamespace):

    def __init__(self, recorder: Recorder, logger: Logger):
        super().__init__(NAMESPACE)
        self._recorder = recorder
        self._logger = logger

    def on_StartRecording(self, _sid) -> None:
        self._logger.info('Received request to start recording')
        try:
            self._recorder.start()
            self._logger.info('Recording was started successfully')
        except Exception as e:
            self._logger.error('Failed to start recording')

    def on_StopRecording(self, _sid) -> None:
        try:
            self._recorder.stop()
            self._logger.info('Recording was stopped successfully')
        except Exception as e:
            self._logger.error('Failed to stop recording')

    async def on_QueryRecordingState(self, _sid):
        is_recording = self._recorder.get_is_recording()
        await self.emit('RecordingState', {'isRecording': is_recording})
