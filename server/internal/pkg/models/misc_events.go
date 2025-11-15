package models

func NewFileRenamedEvent() Event[any] {
	return Event[any]{
		EventId: FileRenamedEvent,
	}
}
