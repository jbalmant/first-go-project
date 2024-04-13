package entity

type EventType string

const (
	AddUser EventType = "AddUser"
	End     EventType = "End"
	Init    EventType = "Init"
	Kill    EventType = "Kill"
)

type Event struct {
	Type    EventType
	Details interface{}
}

func NewAction(eventType EventType) *Event {
	return &Event{
		Type: eventType,
	}
}

type AddUserEvent struct {
	Username string
}

func NewAddUserEvent(username string) *Event {
	return &Event{
		Type: AddUser,
		Details: &AddUserEvent{
			Username: username,
		},
	}
}

type EndGameEvent struct {
}

func NewEndGameEvent() *Event {
	return &Event{Type: End}
}

type InitGameEvent struct {
}

func NewInitGameEvent() *Event {
	return &Event{Type: Init}
}

type KillEvent struct {
	Killer       string
	Victmin      string
	CauseOfDeath string
}

func NewKillEvent(killer string, victmin string, causeOfDeath string) *Event {
	return &Event{
		Type: Kill,
		Details: &KillEvent{
			Killer:       killer,
			Victmin:      victmin,
			CauseOfDeath: causeOfDeath,
		},
	}
}
