import logging
import socketio
import tornado
import threading
from tornado.ioloop import IOLoop
from tkinter import Label, Button, Tk, Frame, Canvas, ttk
import tkinter as tk

from config_loader import ConfigLoader, DATA_DIR_PATH
from recorder import Recorder

sio = socketio.AsyncServer(async_mode='tornado')
logging.basicConfig(level=logging.INFO,
                    format='%(asctime)s - %(levelname)s - %(message)s')
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
        await sio.emit('RecordingStateChange', {'recording': True})
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
        await sio.emit('RecordingStateChange', {'recording': False})
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
    await sio.emit('RecordingState', {'isRecording': is_recording})


def createUI():
    root = Tk()
    root.state('zoomed')

    # frame for the label and buttons
    frame = Frame(root)
    frame.place(relx=0.5, rely=0.5, anchor="c")  # put at center of window

    # frame for the two buttons
    frame = Frame(root)
    # frame.pack(expand=True)
    frame.pack(expand=True)

    recording_state_label = Label(frame, text="Recording", font=("Arial", 100), background='green', padx=100, pady=100)
    recordings_list_label = Label(
        frame, text="List of recordings", font=("Arial", 40))

    frame.grid_rowconfigure(0)

    recording_state_label.grid(column=2, row=0, padx=150)
    recordings_list_label.grid(row=0, column=0, sticky="SW")
    treeview = ttk.Treeview(frame, columns=("size", "lastmod"), height=40)
    treeview.heading("#0", text="File")
    treeview.heading("size", text="Size")
    treeview.heading("lastmod", text="Last modification")
    treeview.insert(
        "",
        tk.END,
        text="README.txt",
        values=("850 bytes", "18:30")
    )
    treeview.grid(column=0, row=1)

    print('grid-size: ', frame.grid_size())
    root.mainloop()
    # TODO: Disconnect SocketIO clients
    server.stop()
    IOLoop.current().stop()
    logger.info("stopped")


thread = threading.Thread(target=createUI)
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
