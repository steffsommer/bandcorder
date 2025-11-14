import { AnimatePresence, motion } from "motion/react";
import { useState } from "react";
import { createBrowserRouter, RouterProvider, Outlet, useLocation } from "react-router-dom";
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
  const location = useLocation();
  return (
    <div className="layout">
      <Header onSettingsClick={() => setShowSettings(true)} />
      <motion.div
        className="full-height-flex-child"
        key={location.pathname}
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.2 }}
      >
        <Outlet />
      </motion.div>
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
