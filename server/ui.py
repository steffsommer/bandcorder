import tkinter as tk
import recording_state_label
import recordings_treeview


class UserInterface(tk.Tk):
    """ Tkinter User interface root

     Displays some information like a list of the previous recordings
     of the current day and the current recording state
    """
    def __init__(self):
        super().__init__()
        self.state('zoomed')
        self._root_frame = tk.Frame(self)
        # component definitions
        self._recording_label = recording_state_label.RecordingStateLabel(self._root_frame)
        self._recodings_tree_view =  recordings_treeview.RecordingsTreeView(self._root_frame)

        # layout
        self._root_frame.grid_rowconfigure(0)

        # component placement
        self._root_frame.place(relx=0.5, rely=0.5, anchor="c")
        self._root_frame.pack(expand=True)
        self._recording_label.grid(column=2, row=0, padx=150)
        self._recodings_tree_view.grid(column=0, row=1)
