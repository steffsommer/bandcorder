import logging
import eventlet
import socketio

from recorder import Recorder

recorder = Recorder()
sio = socketio.Server()
app = socketio.WSGIApp(sio)


logger = logging.getLogger(__name__)

 #@sio.on('connect')
#def test_connect():
#   logger.info('Client connected')

#@sio.on('disconnect')
#def test_disconnect():
#    logger.info('Client disconnected')

@sio.on('StartRecording')
def start_recording(event):
    logger.info('Received request to start recording')
    try:
        recorder.start()
        sio.emit('RecordingStateChange', { 'recording': True })
        logger.info('Recording was started successfully')
    except Exception as e:
        logger.error('Failed to start recording')
        sio.emit('RecordingStateChangeFailed', {
            'requestedState': True,
            'reason': str(e)
        })

@sio.on('StopRecording')
def stop_recording(event):
    try:
        recorder.stop()
        sio.emit('RecordingStateChange', { 'recording': False })
        logger.info('Recording was stopped successfully')
    except Exception as e:
        logger.error('Failed to stop recording')
        sio.emit('RecordingStateChangeFailed', {
            'requestedState': False,
            'reason': str(e)
        })

@sio.on('QueryRecordingState')
def query_recording_state():
    is_recording = recorder.get_is_recording()
    sio.emit('RecordingState', { 'isRecording': is_recording })

if __name__ == '__main__':
    eventlet.wsgi.server(eventlet.listen(('', 5000)), app)