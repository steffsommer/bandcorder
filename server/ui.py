import tkinter as tk
from recording_state_label import RecordingStateLabel
from recordings_treeview import RecordingsTreeView
from active_mic_label import ActiveMicLabel
from active_recording_info import ActiveRecordingInfo, RecordingState
from recording_state_notifier import RecordingStateNotifier
from file_storage_service import FileStorageService


class UserInterface(tk.Tk):
    """ Tkinter User interface root

     Displays some information like a list of the previous recordings
     of the current day and the current recording state
    """

    def __init__(
            self,
            storage_service: FileStorageService,
            notifier: RecordingState
    ):
        super().__init__()
        self.state('zoomed')
        self._root_frame = tk.Frame(self)
        # component definitions
        self._active_mic_label = ActiveMicLabel(self._root_frame, "DUMMY MIC")
        self._recording_label = RecordingStateLabel(self._root_frame, notifier)
        self._recordings_tree_view = RecordingsTreeView(self._root_frame, notifier, storage_service)
        self._active_recording_info = ActiveRecordingInfo(self._root_frame)

        # layout
        self._root_frame.grid_rowconfigure(0)

        # component placement
        self._root_frame.place(relx=0.5, rely=0.5, anchor="c")
        self._active_mic_label.grid(row=0, column=0, sticky='NW')
        self._recording_label.grid(row=0, column=1, padx=150)
        self._recordings_tree_view.grid(row=1, column=0)
        self._active_recording_info.grid(row=1, column=1, padx=150, pady=40, sticky='NW')

        # debug/test code
        state = RecordingState(file_name="dummy_file.wav", seconds=14)
        self._active_recording_info.update(state)
