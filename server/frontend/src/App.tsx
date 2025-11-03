import { useState } from "react";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import "./App.css";
import { Card } from "./components/card/card";
import { Header } from "./components/header/header";
import { IP } from "./components/ip/ip";
import { Recorder } from "./components/recorder/recorder";
import { RecordingsList } from "./components/recordings-list/recordings-list";
import { SettingsModal } from "./components/settings-modal/settings-modal";
import MetronomePage from "./pages/metronome-page";

const router = createBrowserRouter([
  {
    path: "/",
    element: <RecordingPage />,
  },
  {
    path: "/metronome",
    element: <MetronomePage />,
  },
]);

function App() {
  const [showSettings, setShowSettings] = useState(false);
  return (
    <div id="app">
      <Header onSettingsClick={() => setShowSettings(true)} />
      <div className="container">
        <RouterProvider router={router} />;
        <SettingsModal show={showSettings} onClose={() => setShowSettings(false)} />
        <ToastContainer autoClose={2000} />
      </div>
    </div>
  );
}

function RecordingPage() {
  return (
    <div className="recording-layout">
      <Card>
        <RecordingsList />
      </Card>
      <Card>
        <Recorder />
      </Card>
    </div>
  );
}

export default App;
