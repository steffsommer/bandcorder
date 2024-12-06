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

socketio_server = socketio.AsyncServer(async_mode='tornado')
notifier = RecordingStateNotifier()
client_notifier = WebSocketClientNotifier(logger, socketio_server)
notifier.register_subscriber(client_notifier)


data_dir = config[DATA_DIR_PATH]
recorder = Recorder(notifier, data_dir)

recording_controller = RecordingController(recorder, logger)
socketio_server.register_namespace(recording_controller)

def shutdown():
    client_notifier.stop()
    server.stop()
    IOLoop.current().stop()


def run_ui():
    root_widget = ui.UserInterface(recorder, notifier)
    root_widget.mainloop()
    shutdown()


thread = threading.Thread(target=run_ui)
thread.start()

app = tornado.web.Application(
    [
        (r"/socket.io/", socketio.get_tornado_handler(socketio_server)),
    ],
)
server = app.listen(5000)

eventLoopThread = threading.Thread(target=IOLoop.current().start)
eventLoopThread.daemon = True
eventLoopThread.start()

client_notifier.start()
