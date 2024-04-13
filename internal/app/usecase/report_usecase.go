package usecase

import (
	"MyFirstModule/internal/app/entity"
)

type Report struct {
	Name    string
	GameID  int
	Details interface{}
}

type ReportMatchDetail struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

type ReportUsecase interface {
	Generate(game entity.Game) Report
	Deliver(report Report)
}

type gameDetailsReportUsecase struct {
	reportDelivery ReportDelivery
}

func NewGameDetailsReportUsecase(reportDelivery ReportDelivery) *gameDetailsReportUsecase {
	return &gameDetailsReportUsecase{
		reportDelivery: reportDelivery,
	}
}

func (report *gameDetailsReportUsecase) Generate(game entity.Game) Report {
	players := make([]string, 0)
	totalKills := 0
	kills := make(map[string]int, 0)

	for _, player := range game.ListPlayers() {
		players = append(players, player.Name)
		kills[player.Name] = 0
	}

	for _, kill := range game.ListKills() {
		totalKills++

		if kill.Killer == "<world>" {
			kills[kill.Victmin]-- // ? It's OK to negativate the player kills or should stay on zero?
		} else {
			kills[kill.Killer]++
		}
	}

	result := &Report{
		Name:   "GameMatchDetails",
		GameID: game.ID,
		Details: &ReportMatchDetail{
			TotalKills: totalKills,
			Players:    players,
			Kills:      kills,
		},
	}

	return *result
}

func (detailReport *gameDetailsReportUsecase) Deliver(report Report) {
	detailReport.reportDelivery.Deliver(report)
}

type CauseOfDeathDetail struct {
	Kills map[string]int `json:"kill_by_means"`
}

type causeOfDeathReportUsecase struct {
	reportDelivery ReportDelivery
}

func NewCauseOfDeathReportUsecase(reportDelivery ReportDelivery) *causeOfDeathReportUsecase {
	return &causeOfDeathReportUsecase{
		reportDelivery: reportDelivery,
	}
}

func (report *causeOfDeathReportUsecase) Generate(game entity.Game) Report {
	killsByMean := make(map[string]int, 0)

	for _, kill := range game.ListKills() {
		killsByMean[kill.CauseOfDeath]++
	}

	result := &Report{
		Name:   "CauseOfDeath",
		GameID: game.ID,
		Details: &CauseOfDeathDetail{
			Kills: killsByMean,
		},
	}

	return *result
}

func (detailReport *causeOfDeathReportUsecase) Deliver(report Report) {
	detailReport.reportDelivery.Deliver(report)
}
