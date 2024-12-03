from recording_state_notifier import RecordingStateNotifier, RecordingState
import socketio
import asyncio
import logging
import time

INTERAL_SECONDS = 1


class WebSocketClientNotifier(RecordingStateNotifier):

    def __init__(self, logger: logging.Logger, sio: socketio.AsyncServer):
        self._sio = sio
        self._logger = logger
        self._state = RecordingState(
            is_recording=False, duration=0, file_name='')
        self._running = True
        
    def start(self) -> None:
        loop = asyncio.get_event_loop()
        loop.create_task(self._send_periodically())

    def _send_periodically(self) -> None:
        while self._running:
            time.sleep(INTERAL_SECONDS)
            self.on_state_change(self._state)
            self._logger.info(f'Published state to client {self._state}')
            

    def stop(self):
        self._running = False

    def on_state_change(self, state: RecordingState) -> None:
        self._state = state
        if state is not None:
            loop = asyncio.get_event_loop()
            dto = state.toSerializableDict()
            loop.create_task(self._sio.emit('RecordingStateChange', dto))
