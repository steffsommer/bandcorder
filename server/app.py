import logging

from flask import Flask
from flask_socketio import SocketIO, emit
from recorder import Recorder

recorder = Recorder()
app = Flask(__name__)
socketio = SocketIO(app)
logger = logging.getLogger(__name__)

@socketio.on('connect')
def test_connect():
    logger.info('Client connected')

@socketio.on('disconnect')
def test_disconnect():
    logger.info('Client disconnected')

@socketio.on('StartRecording')
def start_recording():
    logger.info('Received request to start recording')
    try:
        recorder.start()
        emit('RecordingStateChange', { 'recording': True }, broadcast=True)
        logger.info('Recording was started successfully')
    except Exception as e:
        logger.error('Failed to start recording')
        emit('RecordingStateChangeFailed', {
            'requestedState': True,
            'reason': str(e)
        }, broadcast=True)

@socketio.on('StopRecording')
def stop_recording():
    try:
        recorder.stop()
        emit('RecordingStateChange', { 'recording': False }, broadcast=True)
        logger.info('Recording was stopped successfully')
    except Exception as e:
        logger.error('Failed to stop recording')
        emit('RecordingStateChangeFailed', {
            'requestedState': False,
            'reason': str(e)
        }, broadcast=True)

@socketio.on('QueryRecordingState')
def query_recording_state():
    is_recording = recorder.get_is_recording()
    emit('RecordingState', { 'isRecording': is_recording })

if __name__ == '__main__':
    socketio.run(app)