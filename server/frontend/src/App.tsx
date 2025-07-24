import "./App.css";
import { Card } from "./components/card/card";
import { Header } from "./components/header/header";

function App() {
  return (
    <div id="app">
      <Header />
      <div className="container">
        <Card>
          <span>Recording list</span>
        </Card>
        <Card>
          <span>Recorder</span>
        </Card>
      </div>
    </div>
  );
}

export default App;
