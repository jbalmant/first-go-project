package main

import (
	"fmt"
	"time"

	"MyFirstModule/internal/app/infrastructure"
	"MyFirstModule/internal/app/usecase"
	"MyFirstModule/pkg"
)

func main() {
	startTime := time.Now()

	logger := pkg.NewLogger(pkg.INFO)

	eventBus := infrastructure.NewEventBus(*logger)

	fileReaderEventDispatcher := usecase.NewFileReaderEventDispatcher(*logger, *eventBus, "./assets/qgames.log")
	quake3Parser := usecase.NewQuake3Parser(eventBus)
	gameDetailReport := usecase.NewGameDetailsReportUsecase(usecase.NewFileReportDelivery())
	causeOfDeathReport := usecase.NewCauseOfDeathReportUsecase(usecase.NewFileReportDelivery())

	gameControlUsecase := usecase.NewGameControlUsecase(*logger, fileReaderEventDispatcher, quake3Parser, []usecase.ReportUsecase{gameDetailReport, causeOfDeathReport})
	gameControlUsecase.Start()

	endTime := time.Now()
	fmt.Println("Elapsed Time: ", endTime.Sub(startTime))
}
