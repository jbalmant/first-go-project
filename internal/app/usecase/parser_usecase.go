package usecase

import (
	"MyFirstModule/internal/app/entity"
	"regexp"
)

type EventType string

const (
	AddUser  = "ClientUserinfoChanged"
	EndGame  = "ShutdownGame"
	InitGame = "InitGame"
	Kill     = "Kill"
)

type parserUsecase interface {
	Parse(line string) *entity.Event
}

type quake3Parser struct{}

func NewQuake3Parser() *quake3Parser {
	return &quake3Parser{}
}

func mapInitGame() *entity.Event {
	return entity.NewInitGameEvent()
}

func mapEndGame() *entity.Event {
	return entity.NewEndGameEvent()
}

func mapLogToAddUserEvent(eventLine string) *entity.Event {
	eventPattern := `n\\([^\\]+)`

	re := regexp.MustCompile(eventPattern)

	submatches := re.FindStringSubmatch(eventLine)

	if len(submatches) < 1 {
		panic("AddUserEvent pattern does not match")
	}

	return entity.NewAddUserEvent(
		submatches[1],
	)
}

func mapLogToKillEvent(eventLine string) *entity.Event {
	eventPattern := `Kill: \d+\s\d+\s\d+:\s(.*?)\skilled\s(.*?)\sby\s(\S+)`

	re := regexp.MustCompile(eventPattern)

	submatches := re.FindStringSubmatch(eventLine)

	if len(submatches) < 4 {
		panic("KillEvent pattern does not match")
	}

	return entity.NewKillEvent(
		submatches[1],
		submatches[2],
		submatches[3],
	)
}

func (parser *quake3Parser) Parse(line string) *entity.Event {
	pattern := `^\s*\d{1,3}:\d{2}.*(InitGame|ShutdownGame|Kill|ClientUserinfoChanged)`

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(line)

	if match == nil {
		return nil
	}

	action := EventType(match[1])

	switch action {
	case AddUser:
		return mapLogToAddUserEvent(line)
	case EndGame:
		return mapEndGame()
	case InitGame:
		return mapInitGame()
	case Kill:
		return mapLogToKillEvent(line)
	default:
		panic("Event mapper not found")
	}
}
