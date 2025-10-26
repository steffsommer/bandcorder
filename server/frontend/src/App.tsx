import { useState } from "react";
import { ToastContainer } from "react-toastify";
import "./App.css";
import { Card } from "./components/card/card";
import { Header } from "./components/header/header";
import { IP } from "./components/ip/ip";
import { Recorder } from "./components/recorder/recorder";
import { RecordingsList } from "./components/recordings-list/recordings-list";
import { SettingsModal } from "./components/settings-modal/settings-modal";

function App() {
  const [showSettings, setShowSettings] = useState(false);
  return (
    <div id="app">
      <Header onSettingsClick={() => setShowSettings(true)} />
      <div className="container">
        <div className="recording-layout">
          <Card>
            <RecordingsList />
          </Card>
          <Card>
            <Recorder />
          </Card>
        </div>
        <IP className="ip" />
      </div>
      <SettingsModal show={showSettings} onClose={() => setShowSettings(false)} />
      <ToastContainer autoClose={2000} />
    </div>
  );
}

export default App;
