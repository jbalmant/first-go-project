package main

import (
	"MyFirstModule/internal/app/usecase"
	"MyFirstModule/pkg"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	logger := pkg.NewLogger(pkg.INFO)

	fileReaderEventDispatcher := usecase.NewFileReaderEventDispatcher(*logger, "./assets/qgames.log")
	quake3Parser := usecase.NewQuake3Parser()
	gameDetailReport := usecase.NewGameDetailsReportUsecase(usecase.NewFileReportDelivery())
	causeOfDeathReport := usecase.NewCauseOfDeathReportUsecase(usecase.NewFileReportDelivery())

	gameControlUsecase := usecase.NewGameControlUsecase(*logger, fileReaderEventDispatcher, quake3Parser, []usecase.ReportUsecase{gameDetailReport, causeOfDeathReport})
	gameControlUsecase.Start()

	endTime := time.Now()
	fmt.Println("Elapsed Time: ", endTime.Sub(startTime))
}
