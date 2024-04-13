package usecase

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ReportDelivery interface {
	Deliver(report Report)
}

type fileReportDelivery struct {
	timestamp time.Time
}

func NewFileReportDelivery() *fileReportDelivery {
	return &fileReportDelivery{
		timestamp: time.Now(),
	}
}

func (fr *fileReportDelivery) Deliver(report Report) {
	file, err := os.OpenFile(fmt.Sprintf("%v-%v.txt", report.Name, fr.timestamp.Format("2006_01_02_15_04_05")), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o644)
	if err != nil {
		panic("Deliver - An error occurred when opening the file")
	}
	defer file.Close()

	jsonData, _ := json.Marshal(report.Details)

	key := fmt.Sprintf("game_%d", report.GameID)
	dataMap := map[string]interface{}{
		key: json.RawMessage(jsonData),
	}

	dataToPrint, _ := json.MarshalIndent(dataMap, "", "  ")

	_, err = file.WriteString(string(dataToPrint))
	if err != nil {
		panic("Deliver - An error occurred when writing on file")
	}
}
