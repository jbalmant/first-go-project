package usecase

import (
	"MyFirstModule/pkg"
	"bufio"
	"os"
)

type EventListener func(string)

type eventDispatcherUsecase interface {
	SetEventHandler(eventHandler EventListener)
	Start()
}

type fileReaderEventDispatcher struct {
	logger       pkg.Logger
	eventHandler EventListener
	filePath     string
}

func NewFileReaderEventDispatcher(logger pkg.Logger, filePath string) *fileReaderEventDispatcher {
	return &fileReaderEventDispatcher{
		logger:       logger,
		eventHandler: nil,
		filePath:     filePath,
	}
}

func (fileReader *fileReaderEventDispatcher) SetEventHandler(eventHandler EventListener) {
	fileReader.eventHandler = eventHandler
}

func (fileReader *fileReaderEventDispatcher) Start() {
	file, err := os.Open(fileReader.filePath)
	if err != nil {
		fileReader.logger.Error(err.Error())
		panic("An error occurred when opening file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileReader.eventHandler(line)
	}
}
