import tkinter as tk
from tkinter import ttk


class RecordingsTreeView(ttk.Treeview):
    def __init__(self, parent):
        super().__init__(parent, columns=("size", "lastmod"), height=40)
        self.heading("#0", text="File")
        self.heading("size", text="Size")
        self.heading("lastmod", text="Last modification")
        self.insert(
            "",
            tk.END,
            text="README.txt",
            values=("850 bytes", "18:30")
        )
