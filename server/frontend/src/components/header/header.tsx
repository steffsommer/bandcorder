import { Card } from "../card/card";
import "./header.css";
import logo from "../../assets/images/logo.svg";
import { FiSettings } from "react-icons/fi";
import { Button } from "../button/button";
import { PiMicrophoneStageBold } from "react-icons/pi";
import { TbMetronome } from "react-icons/tb";

interface Props {
  onSettingsClick: () => void;
}

export function Header({ onSettingsClick }: Props) {
  const currentPath = window.location.pathname;

  return (
    <header>
      <Card className="logo">
        <a href="/" className="logo">
          <img className="logo-img" src={logo} />
          <h1>BANDCORDER</h1>
        </a>
      </Card>
      <div className="nav-items">
        <a href="/" className={`nav-item ${currentPath === "/" ? "active" : ""}`}>
          <PiMicrophoneStageBold size="1.3em" />
          <span>RECORD</span>
        </a>
        <a href="/metronome" className={`nav-item ${currentPath === "/metronome" ? "active" : ""}`}>
          <TbMetronome size="1.5em" />
          <span>METRONOME</span>
        </a>
      </div>
      <Button className="settings-btn" onClick={onSettingsClick}>
        <FiSettings size="2em" />
      </Button>
    </header>
  );
}
