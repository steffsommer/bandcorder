from recording_state_notifier import RecordingStateNotifier, RecordingState
import socketio
import asyncio
from typing import cast


class WebSocketClientNotifier(RecordingStateNotifier):

    def __init__(self, sio: socketio.AsyncServer):
        self._sio = sio

    def on_state_change(self, state: RecordingState) -> None:
        loop = asyncio.get_event_loop()
        dto = state.toSerializableDict()
        loop.create_task(self._sio.emit('RecordingStateChange', dto))
