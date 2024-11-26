# Bandcorder

Disclaimer: This is early stage software

## About

Simply put, Bandcorder allows to create, start & stop audio recordings easily in a multi-client, shared session using a smartphone app.
The recordings may be done using a (high quality) microphone attached to a server (potentially a laptop).

## Motivation

During jam sessions or Band practice the need to do simple recordings arises frequently. Since rehearsal rooms
are usually not equipped with the gear to perform this type of recording, musicians often resort to
using their phones. This comes with the following problems:

- At best mediocore audio quality with little to no configuration options
- Having to manually share the recordings with bandmates via different channels (e.g. messengers) later on

Bandcorder solves both of these problems in a simple manner.

## How it works

When the client app gets opened and is connected to a Wifi, it scans the network for a Bandcorder server.
As soon as the server is found, a connection is established and the client is part of a recording session.
Any client may start or stop recordings and the recording state is synced between all participants.  
The server uses the default Microphone of the system to create the recordings. The recording files are
saved in a directory corresponding to the current day with a file name resembling the current time.
To share the recordings quickly with bandmates, a Cloud storage sync (e.g. using Dropbox) may be set
for the folder.

## Configuration

Copy `config.yml.template` to `config.yml` and specify a valid data directory.

## Development setup

### Requirements

- python3
- Flutter >= 3.18

### Server setup
```bash
    # (optional) Create venv
    python -m venv .venv && source ./.venv/Scripts/activate
    # Install dependencies
    cd server && pip install -e .
```

## Todos for 1.0

- [ ] Mobile UI (Flutter)
- [x] Reconfiguration as yml file
  - [x] Specify recording directory
- [ ] Desktop UI Tkinter
- [x] Replace Flask with socket.io server
- [x] Create virtualenv
- [ ] Create Release.exe
- [x] Put recordings in date folders
- [ ] Create a fancy logo

## Accepted Limitations for 1.0

- Only one server instance per Wifi network
- Plaintext communication (no secrets are transmitted though)
- Only tested on Windows 11
