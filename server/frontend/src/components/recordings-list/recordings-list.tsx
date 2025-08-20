import {
  Recording,
  RecordingsListEntry,
} from "./recordings-list-entry/recordings-list-entry";
import "./recordings-list.css";
import React from "react";

interface Props {
  recordings: Recording[];
}

export const RecordingsList: React.FC<Props> = ({ recordings }) => {
  return (
    <div className="recordings-list">
      <h2 className="descriptive-header">Todays' recordings</h2>
      <div className="recordings">
        {recordings.length === 0 ? (
          <p>No items to display</p>
        ) : (
          recordings.map((item, index) => (
            <RecordingsListEntry recording={item} key={index} />
          ))
        )}
      </div>
    </div>
  );
};
