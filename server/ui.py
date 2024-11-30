import tkinter as tk
from tkinter import ttk
import recording_state_label

# Tkinter top level widget implementation
# Displays some information like a list of the previous recordings
# of the current day and the current recording state


class UserInterface(tk.Tk):
    def __init__(self):
        super().__init__()
        self.state('zoomed')

        # frame for the label and buttons
        frame = tk.Frame(self)
        frame.place(relx=0.5, rely=0.5, anchor="c")  # put at center of window

        # frame for the two buttons
        frame = tk.Frame(self)
        # frame.pack(expand=True)
        frame.pack(expand=True)
        
        recording_label = recording_state_label.RecordingStateLabel(frame)

        # recording_state_label = tk.Label(frame, text="Recording", font=(
        #     "Arial", 100), background='green', padx=100, pady=100)
        recordings_list_label = tk.Label(
            frame, text="List of recordings", font=("Arial", 40))

        frame.grid_rowconfigure(0)

        recording_label.grid(column=2, row=0, padx=150)
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
