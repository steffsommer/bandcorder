import tkinter as tk
from tkinter import ttk


class RecordingsTreeView(tk.Frame):
  
    def __init__(self, parent):
        super().__init__(parent)

        # component definitions
        self._treeview = self._get_configured_treeview()
        self._recordings_list_label = tk.Label(
            self,
            text="List of recordings",
            font=("Arial", 40)
        )

        # placement
        self._treeview.grid(row=1, column=0)
        self._recordings_list_label.grid(row=0, column=0)


    def _get_configured_treeview(self) -> ttk.Treeview:
        rec_list = ttk.Treeview(self, columns=("size", "lastmod"), height=20)
        rec_list.heading("#0", text="File")
        rec_list.heading("size", text="Size")
        rec_list.heading("lastmod", text="Last modification")
        rec_list.insert(
            "",
            tk.END,
            text="README.txt",
            values=("850 bytes", "18:30")
        )
        return rec_list