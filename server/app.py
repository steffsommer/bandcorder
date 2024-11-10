from flask import Flask
from recorder import Recorder

app = Flask(__name__)
recorder = Recorder()

@app.route("/recording/state", methods=["GET"])
def get_recordings():
    raise "not yet implemented"

@app.route("/recording/start", methods=["POST"])
def start_recording():
    recorder.start()

@app.route("/recording/stop", methods=["POST"])
def stop_recording():
    raise "not yet implemented"
