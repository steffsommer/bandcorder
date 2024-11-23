import logging
import eventlet
import socketio
from config_loader import ConfigLoader, DATA_DIR_PATH
from recorder import Recorder

sio = socketio.Server()
app = socketio.WSGIApp(sio)

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)

config_loader = ConfigLoader(logger)
config = config_loader.load_config()

data_dir = config[DATA_DIR_PATH]
recorder = Recorder(data_dir)

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
    eventlet.wsgi.server(eventlet.listen(('', 5000)), app, log=logger)