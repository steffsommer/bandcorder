from recording_state_notifier import RecordingStateNotifier, RecordingState
import socketio


class WebSocketClientNotifier(RecordingStateNotifier):

    def __init__(self, sio: socketio.AsyncServer):
        self._sio = sio

    async def on_state_change(self, state: RecordingState) -> None:
        await self._sio.emit(
            'RecordingStateChange',
            {
                'isRecording': state.is_recording,
                'fileName': state.file_name,
                'duration': state.duration,
            }
        )
