#!/usr/bin/env python

import threading
import socketio
from socketio.exceptions import TimeoutError
import sys

CMD_SEPARATOR = "/"

sio = socketio.SimpleClient()
sio.connect('http://localhost:5000')
print("connected")

running = True
def receive():
    while running:
        try:
            received_event = sio.receive(1)
            print('Received received_event: ', received_event)
        except TimeoutError:
            pass
        except Exception as e:
            print(str(e))
            sys.exit()

receiver_thread = threading.Thread(target=receive)
receiver_thread.start()

print('messages are received through a background thread')
print(f'To send enter "s{CMD_SEPARATOR}<event>{CMD_SEPARATOR}<payload>". Enter "q" to quit.\n')
user_input = ""
while user_input != "q":
    user_input = input('$ ')
    if user_input.startswith("s"):
        if user_input.count(CMD_SEPARATOR) == 1:
            (_, event) = user_input.split(CMD_SEPARATOR)
            sio.emit(event)
        elif user_input.count(CMD_SEPARATOR) == 2:
            (_, event, event_data) = user_input.split(CMD_SEPARATOR)
            sio.emit(event, event_data)
        else:
            print("Invalid syntax. Skipping command.\n")
            continue
        print("event sent successfully!")
    else:
        print("Invalid syntax. Skipping command.\n")