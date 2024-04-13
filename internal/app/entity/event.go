package entity

type EventType string

const (
	AddPlayer EventType = "AddUser"
	End       EventType = "End"
	Init      EventType = "Init"
	Kill      EventType = "Kill"
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

type AddPlayerEvent struct {
	Username string
}

func NewAddPlayerEvent(username string) *Event {
	return &Event{
		Type: AddPlayer,
		Details: &AddPlayerEvent{
			Username: username,
		},
	}
}

type EndGameEvent struct{}

func NewEndGameEvent() *Event {
	return &Event{Type: End}
}

type InitGameEvent struct{}

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
