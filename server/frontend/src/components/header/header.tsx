import { Card } from "../card/card";
import "./header.css";
import logo from "../../assets/images/logo.svg";
import { FiSettings } from "react-icons/fi";
import { Button } from "../button/button";
import { PiMicrophoneStageBold } from "react-icons/pi";
import { TbMetronome } from "react-icons/tb";
import { Link } from "react-router-dom";

interface Props {
  onSettingsClick: () => void;
}

export function Header({ onSettingsClick }: Props) {
  const currentPath = window.location.pathname;

  return (
    <header>
      <Card className="logo">
        <Link to="/" className="logo">
          <img className="logo-img" src={logo} />
          <h1>BANDCORDER</h1>
        </Link>
      </Card>
      <div className="nav-items">
        <Link to="/" className={`nav-item ${currentPath === "/" ? "active" : ""}`}>
          <PiMicrophoneStageBold size="1.3em" />
          <span>RECORD</span>
        </Link>
        <Link
          to="/metronome"
          className={`nav-item ${currentPath === "/metronome" ? "active" : ""}`}
        >
          <TbMetronome size="1.5em" />
          <span>METRONOME</span>
        </Link>
      </div>
      <Button className="settings-btn" onClick={onSettingsClick}>
        <FiSettings size="2em" />
      </Button>
    </header>
  );
}
