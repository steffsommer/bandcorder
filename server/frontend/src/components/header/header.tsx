import { Card } from "../card/card";
import "./header.css";
import logo from "../../assets/images/logo.svg";
import { FiSettings } from "react-icons/fi";
import { Button } from "../button/button";

export function Header() {
  return (
    <header>
      <Card className="logo">
        <img className="logo-img" src={logo} />
        <h1>BANDCORDER</h1>
      </Card>
      <Button className="settings-btn">
        <FiSettings size="2em" />
      </Button>
    </header>
  );
}
