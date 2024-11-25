import logging
import socketio
import tornado

from config_loader import ConfigLoader, DATA_DIR_PATH
from recorder import Recorder

sio = socketio.AsyncServer(async_mode='tornado')
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)

config_loader = ConfigLoader(logger)
config = config_loader.load_config()

data_dir = config[DATA_DIR_PATH]
recorder = Recorder(data_dir)

@sio.on('StartRecording')
async def start_recording(event):
    logger.info('Received request to start recording')
    try:
        recorder.start()
        await sio.emit('RecordingStateChange', { 'recording': True })
        logger.info('Recording was started successfully')
    except Exception as e:
        logger.error('Failed to start recording')
        await sio.emit('RecordingStateChangeFailed', {
            'requestedState': True,
            'reason': str(e)
        })

@sio.on('StopRecording')
async def stop_recording(event):
    try:
        recorder.stop()
        await sio.emit('RecordingStateChange', { 'recording': False })
        logger.info('Recording was stopped successfully')
    except Exception as e:
        logger.error('Failed to stop recording')
        await sio.emit('RecordingStateChangeFailed', {
            'requestedState': False,
            'reason': str(e)
        })

@sio.on('QueryRecordingState')
async def query_recording_state():
    is_recording = recorder.get_is_recording()
    await sio.emit('RecordingState', { 'isRecording': is_recording })
    
if __name__ == '__main__':
    app = tornado.web.Application(
        [
            (r"/socket.io/", socketio.get_tornado_handler(sio)),
        ],
    )
    app.listen(5000)
    tornado.ioloop.IOLoop.current().start()
