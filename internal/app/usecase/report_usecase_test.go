package usecase_test

import (
	"testing"

	"MyFirstModule/internal/app/entity"
	"MyFirstModule/internal/app/usecase"
	"MyFirstModule/test_utils"
)

func SetupBasicGame() *entity.Game {
	game := entity.NewGame(1)
	game.AddPlayer(*entity.NewPlayer("Jonaba"))
	game.AddPlayer(*entity.NewPlayer("Yan"))
	game.AddKill(*&entity.KillEvent{"Jonaba", "Yan", "MOD_TRIGGER_HURT"})
	return game
}

func TestReportUsecaseShouldGenerateWithCorrectNameAndGameId(t *testing.T) {
	// Arrange
	report := usecase.NewGameDetailsReportUsecase(nil)
	game := SetupBasicGame()
	expectedReportName := "GameMatchDetails"
	expectedGameId := 1

	// Act
	result := report.Generate(*game)

	// Assert
	test_utils.AssertEquals(t, expectedReportName, result.Name)
	test_utils.AssertEquals(t, expectedGameId, result.GameID)
}

func TestReportUsecaseShouldGenerateWithCorrectPlayers(t *testing.T) {
	// Arrange
	report := usecase.NewGameDetailsReportUsecase(nil)
	game := SetupBasicGame()

	// Act
	result := report.Generate(*game)

	// Assert
	details := result.Details.(*usecase.ReportMatchDetail)
	test_utils.AssertSlicesEquals(t, details.Players, []string{"Jonaba", "Yan"})
}

func TestReportUsecaseShouldGenerateWithCorrectRanks(t *testing.T) {
	// Arrange
	report := usecase.NewGameDetailsReportUsecase(nil)
	game := SetupBasicGame()

	// Act
	result := report.Generate(*game)

	// Assert
	details := result.Details.(*usecase.ReportMatchDetail)
	test_utils.AssertMapEquals(t, details.Ranks, map[string]int{"Jonaba": 10, "Yan": -5})
}

func TestReportUsecaseShouldGenerateWithCorrectKills(t *testing.T) {
	// Arrange
	report := usecase.NewGameDetailsReportUsecase(nil)
	game := SetupBasicGame()

	// Act
	result := report.Generate(*game)

	// Assert
	details := result.Details.(*usecase.ReportMatchDetail)
	test_utils.AssertMapEquals(t, details.Kills, map[string]int{"Jonaba": 1, "Yan": 0})
}

func TestReportUsecaseShouldGenerateWithCorrectKillsWhenKilledByWorld(t *testing.T) {
	// Arrange
	report := usecase.NewGameDetailsReportUsecase(nil)
	game := SetupBasicGame()
	game.AddKill(*&entity.KillEvent{"<world>", "Yan", "MOD_TRIGGER_HURT"})

	// Act
	result := report.Generate(*game)

	// Assert
	details := result.Details.(*usecase.ReportMatchDetail)
	test_utils.AssertMapEquals(t, details.Kills, map[string]int{"Jonaba": 1, "Yan": -1})
}

func TestReportUsecaseShouldGenerateWithCorrectTotalKills(t *testing.T) {
	// Arrange
	report := usecase.NewGameDetailsReportUsecase(nil)
	game := SetupBasicGame()

	// Act
	result := report.Generate(*game)

	// Assert
	details := result.Details.(*usecase.ReportMatchDetail)
	if details.TotalKills != 1 {
		t.Errorf("Report TotalKills does not match %v", details.TotalKills)
	}
}
