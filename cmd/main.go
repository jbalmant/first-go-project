package main

import (
	"flag"
	"fmt"
	"time"

	"MyFirstModule/internal/app/usecase"
	"MyFirstModule/pkg"
)

func main() {
	startTime := time.Now()
	logger := pkg.NewLogger(pkg.INFO)

	path := flag.String("path", "./assets/qgames.log", "The file path with logs")
	flag.Parse()

	fileReaderEventDispatcher := usecase.NewFileReaderEventDispatcher(*logger, *path)
	quake3Parser := usecase.NewQuake3Parser()
	gameDetailReport := usecase.NewGameDetailsReportUsecase(usecase.NewFileReportDelivery())
	causeOfDeathReport := usecase.NewCauseOfDeathReportUsecase(usecase.NewFileReportDelivery())

	gameControlUsecase := usecase.NewGameControlUsecase(*logger, fileReaderEventDispatcher, quake3Parser, []usecase.ReportUsecase{gameDetailReport, causeOfDeathReport})
	gameControlUsecase.Start()

	endTime := time.Now()
	logger.Info(fmt.Sprintf("Elapsed Time: %v", endTime.Sub(startTime)))
}
