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
        self.frame = tk.Frame(self)
        # child components
        self.recording_label = recording_state_label.RecordingStateLabel(self.frame)
        self.recodings_tree_view =  recordings_treeview.RecordingsTreeView(self.frame)
        self.recordings_list_label = tk.Label(
            self.frame, text="List of recordings", font=("Arial", 40))


        # placements
        self.frame.place(relx=0.5, rely=0.5, anchor="c")
        self.frame.pack(expand=True)
        self.recording_label.grid(column=2, row=0, padx=150)
        

        self.frame.grid_rowconfigure(0)


        self.recordings_list_label.grid(row=0, column=0, sticky="SW")
        self.recodings_tree_view.grid(column=0, row=1)
