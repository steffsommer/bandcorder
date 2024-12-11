from recording_state_notifier import RecordingStateNotifier, RecordingState
import socketio
import asyncio
import logging


class WebSocketClientNotifier(RecordingStateNotifier):

    def __init__(self, logger: logging.Logger, sio: socketio.AsyncServer):
        self._sio = sio
        self._logger = logger
        self.loop_send_task = None

    def stop(self) -> None:
        self.loop_send_task.cancel()

    def on_state_change(self, state: RecordingState) -> None:
        self._state = state
        if state is not None:
            dto = state.toSerializableDict()
            loop = asyncio.get_event_loop()
            loop.create_task(self._sio.emit('RecordingStateChange', dto))
