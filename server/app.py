import logging
import socketio
import tornado
import threading
from tornado.ioloop import IOLoop
from config_loader import ConfigLoader, DATA_DIR_PATH
from recorder import Recorder
from recording_state_notifier import RecordingStateNotifier
from websocket_notifier import WebSocketClientNotifier
from recording_controller import RecordingController
import ui

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

config_loader = ConfigLoader(logger)
config = config_loader.load_config()

sio = socketio.AsyncServer(async_mode='tornado')
notifier = RecordingStateNotifier()
notifier.register_subscriber(WebSocketClientNotifier(sio))

data_dir = config[DATA_DIR_PATH]
recorder = Recorder(notifier, data_dir)

recording_controller = RecordingController(recorder, logger)
sio.register_namespace(recording_controller)


def run_ui():
    root_widget = ui.UserInterface(recorder, notifier)
    root_widget.mainloop()
    server.stop()
    IOLoop.current().stop()


thread = threading.Thread(target=run_ui)
thread.start()

app = tornado.web.Application(
    [
        (r"/socket.io/", socketio.get_tornado_handler(sio)),
    ],
)
server = app.listen(5000)

eventLoopThread = threading.Thread(target=IOLoop.current().start)
eventLoopThread.daemon = True
eventLoopThread.start()
