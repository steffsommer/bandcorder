package models

func NewSettingsUpdatedEvent(s Settings) Event[Settings] {
	return Event[Settings]{
		EventId: SettingsUpdatedEvent,
		Data:    s,
	}
}
