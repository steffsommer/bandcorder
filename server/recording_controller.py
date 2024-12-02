import socketio
from recorder import Recorder
from logging import Logger
from recording_state_notifier import RecordingStateNotifier

NAMESPACE = '/'


class RecordingController(socketio.AsyncNamespace):

    def __init__(
            self,
            recorder: Recorder,
            logger: Logger,
            notifier: RecordingStateNotifier,
    ):
        super().__init__(NAMESPACE)
        self._recorder = recorder
        self._logger = logger
        self._notifier = notifier

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

    async def on_QueryRecordingState(self, sid) -> None:
        state = self._notifier.get_current_state()
        dto = state.toSerializableDict()
        await self.emit('RecordingState', dto, to=sid)
