from recording_state_notifier import RecordingStateNotifier, RecordingState
import socketio
import asyncio
import logging


class WebSocketClientNotifier(RecordingStateNotifier):

    def __init__(self, logger: logging.Logger, sio: socketio.AsyncServer):
        self._sio = sio
        self._logger = logger

    def on_state_change(self, state: RecordingState) -> None:
        self._state = state
        if state is not None:
            dto = state.toSerializableDict()
            asyncio.create_task(self._sio.emit('RecordingStateChange', dto))
