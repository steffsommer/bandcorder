import { useState } from "react";
import { createBrowserRouter, RouterProvider, Outlet } from "react-router-dom";
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
    element: <Layout />,
    children: [
      {
        path: "/",
        element: <RecordingPage />,
      },
      {
        path: "/metronome",
        element: <MetronomePage />,
      },
    ],
  },
]);

function App() {
  return (
    <div id="app">
      <div className="container">
        <RouterProvider router={router} />
        <ToastContainer autoClose={2000} />
      </div>
    </div>
  );
}

function Layout() {
  const [showSettings, setShowSettings] = useState(false);
  return (
    <div className="layout">
      <Header onSettingsClick={() => setShowSettings(true)} />
      <Outlet />
      <SettingsModal show={showSettings} onClose={() => setShowSettings(false)} />
      <IP className="ip" />
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
