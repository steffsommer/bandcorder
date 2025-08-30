import "./App.css";
import { Card } from "./components/card/card";
import { Header } from "./components/header/header";
import { Recorder } from "./components/recorder/recorder";
import { RecordingsList } from "./components/recordings-list/recordings-list";
import { Recording } from "./components/recordings-list/recordings-list-entry/recordings-list-entry";

function App() {
  return (
    <div id="app">
        <Header />
        <div className="container">
          <Card>
            <RecordingsList recordings={recordings} />
          </Card>
          <Card>
            <Recorder />
          </Card>
        </div>
    </div>
  );
}

export default App;

const recordings: Recording[] = [
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
  {
    name: "recording-1",
    duration: "100s",
  },
];

