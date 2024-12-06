from recording_state_notifier import RecordingStateNotifier, RecordingState
import socketio
import asyncio
import logging

INTERVAL_SECONDS = 1


class WebSocketClientNotifier(RecordingStateNotifier):

    def __init__(self, logger: logging.Logger, sio: socketio.AsyncServer):
        self._sio = sio
        self._logger = logger
        self._state = RecordingState(
            is_recording=False, duration=0, file_name='')
        self.loop_send_task = None

    def start(self) -> None:
        loop = asyncio.get_event_loop()
        self.loop_send_task = loop.create_task(self._send_periodically())

    async def _send_periodically(self) -> None:
        while True:
            self.on_state_change(self._state)
            self._logger.info(f'Published state to client {self._state}')
            await asyncio.sleep(INTERVAL_SECONDS)

    def stop(self):
        self.loop_send_task.cancel()

    def on_state_change(self, state: RecordingState) -> None:
        self._state = state
        if state is not None:
            loop = asyncio.get_event_loop()
            dto = state.toSerializableDict()
            loop.create_task(self._sio.emit('RecordingStateChange', dto))
