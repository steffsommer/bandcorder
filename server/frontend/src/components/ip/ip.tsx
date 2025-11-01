import { useEffect, useState } from "react";
import "./ip.css";
import "../../../wailsjs/go/services/IPService";
import { GetLocalIP } from "../../../wailsjs/go/services/IPService";

interface Props {
  className?: string;
}

export function IP({ className }: Props) {
  const [ip, setIp] = useState("");
  useEffect(() => {
    GetLocalIP()
      .then((ip) => {
        setIp(ip);
      })
      .catch(() => {
        setIp("Failed to fetch IP");
      });
  }, []);
  return (
    <span className={(className ?? "") + " ip-container"}>
      <span className="ip-logo">IP</span>
      <span>{ip}</span>
    </span>
  );
}
