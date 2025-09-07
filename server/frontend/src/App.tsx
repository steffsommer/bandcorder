import "./App.css";
import { Card } from "./components/card/card";
import { Header } from "./components/header/header";
import { Recorder } from "./components/recorder/recorder";
import { RecordingsList } from "./components/recordings-list/recordings-list";

function App() {
  return (
    <div id="app">
        <Header />
        <div className="container">
          <Card>
            <RecordingsList />
          </Card>
          <Card>
            <Recorder />
          </Card>
        </div>
    </div>
  );
}

export default App;

