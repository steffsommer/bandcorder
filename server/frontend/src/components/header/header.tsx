import { Card } from "../card/card";
import "./header.css";
import logo from "../../assets/images/logo.svg";
import { FiSettings } from "react-icons/fi";

export function Header() {
  return (
    <header>
      <Card className="logo">
        <img className="logo-img" src={logo} />
        <h1>BANDCORDER</h1>
      </Card>
      <Card className="centered">
        <FiSettings size="2em" />
      </Card>
    </header>
  );
}
