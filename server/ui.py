from tkinter import ttk
from recording_state_label import RecordingStateLabel
from recordings_treeview import RecordingsTreeView
from active_mic_label import ActiveMicLabel
from active_recording_info import ActiveRecordingInfo, RecordingState
from recording_state_notifier import RecordingStateNotifier
from file_storage_service import FileStorageService
import customtkinter as ctk


class UserInterface(ctk.CTk):
    """ Tkinter User interface root

     Displays some information like a list of the previous recordings
     of the current day and the current recording state
    """

    def __init__(
            self,
            storage_service: FileStorageService,
            notifier: RecordingStateNotifier,
    ):
        super().__init__()
        self._customize_styles()
        # hack to start the UI maximized: https://github.com/TomSchimansky/CustomTkinter/discussions/1500
        self._state_before_windows_set_titlebar_color = 'zoomed'
        self._root_frame = ctk.CTkFrame(self)

        # component definitions
        self._active_mic_label = ActiveMicLabel(self._root_frame)
        self._recording_state_label = RecordingStateLabel(self._root_frame, notifier)
        self._recordings_tree_view = RecordingsTreeView(
            self._root_frame, notifier, storage_service)
        self._active_recording_info = ActiveRecordingInfo(
            self._root_frame, notifier)

        # layout
        self._root_frame.grid_rowconfigure(0, weight=1)
        self._root_frame.grid_rowconfigure(1, weight=1)
        self._root_frame.grid_columnconfigure(0, weight=1)
        self._root_frame.grid_columnconfigure(1, weight=1)

        # component placement
        self._root_frame.place(relx=0.5, rely=0.5, anchor="c")
        self._active_mic_label.grid(row=0, column=0)
        self._recording_state_label.grid(row=0, column=1)
        self._recordings_tree_view.grid(row=1, column=0)
        self._active_recording_info.grid(row=1, column=1)

    def _customize_styles(self):
        bg_color = self._apply_appearance_mode(
            ctk.ThemeManager.theme["CTkFrame"]["fg_color"])
        text_color = self._apply_appearance_mode(
            ctk.ThemeManager.theme["CTkLabel"]["text_color"])
        selected_color = self._apply_appearance_mode(
            ctk.ThemeManager.theme["CTkButton"]["fg_color"])
        treestyle = ttk.Style()
        treestyle.theme_use('default')
        treestyle.configure("Treeview", background=bg_color,
                            foreground=text_color, fieldbackground=bg_color, borderwidth=0)
        treestyle.map('Treeview', background=[('selected', bg_color)], foreground=[
                      ('selected', selected_color)])
        self.bind("<<TreeviewSelect>>", lambda event: self.focus_set())
