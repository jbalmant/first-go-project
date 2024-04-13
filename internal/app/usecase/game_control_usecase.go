package usecase

import (
	"MyFirstModule/internal/app/entity"
	"MyFirstModule/pkg"
)

type gameControlUsecase struct {
	logger          pkg.Logger
	eventDispatcher eventDispatcherUsecase
	parser          parserUsecase
	reports         []ReportUsecase
}

func NewGameControlUsecase(
	logger pkg.Logger,
	eventDispatcher eventDispatcherUsecase,
	parser parserUsecase,
	reports []ReportUsecase,

) *gameControlUsecase {
	return &gameControlUsecase{
		logger:          logger,
		eventDispatcher: eventDispatcher,
		parser:          parser,
		reports:         reports,
	}
}

func (control *gameControlUsecase) generateReports(game entity.Game) {
	for _, report := range control.reports {
		result := report.Generate(game)
		report.Deliver(result)
	}
}

func (control *gameControlUsecase) Start() {
	var game *entity.Game = nil
	gameCount := 0

	eventHandler := func(line string) {
		event := control.parser.Parse(line)

		if event == nil {
			return
		}

		switch event.Type {
		case entity.AddUser:
			if game == nil {
				panic("Adding user to a null game")
			}
			detail := event.Details.(*entity.AddUserEvent)
			game.AddPlayer(*entity.NewPlayer(detail.Username))
		case entity.End:
			// TODO: The log shows inconsistencies, new games are initiated without properly closing the previous ones.
			control.generateReports(*game)
			game = nil
		case entity.Init:
			if game != nil {
				control.logger.Error("New Game Without End Previous")
			}
			gameCount++
			game = entity.NewGame(gameCount)
		case entity.Kill:
			game.AddKill(*event.Details.(*entity.KillEvent))
		default:
			panic("Event not mapped")
		}
	}

	control.eventDispatcher.SetEventHandler(eventHandler)
	control.eventDispatcher.Start()
}
