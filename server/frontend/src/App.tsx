import { useState } from "react";
import "./App.css";
import { Card } from "./components/card/card";
import { Header } from "./components/header/header";
import { Recorder } from "./components/recorder/recorder";
import { RecordingsList } from "./components/recordings-list/recordings-list";
import { SettingsDialog } from "./components/settings-sidebar/settings-sidebar";

function App() {
  const [showSettings, setShowSettings] = useState(false);
  return (
    <div id="app">
      <Header onSettingsClick={() => setShowSettings(true)} />
      <div className="container">
        <Card>
          <RecordingsList />
        </Card>
        <Card>
          <Recorder />
        </Card>
      </div>
      <SettingsDialog
        show={showSettings}
        onClose={() => setShowSettings(false)}
      />
    </div>
  );
}

export default App;
