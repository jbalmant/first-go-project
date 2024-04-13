package usecase

import (
	"bufio"
	"os"

	"MyFirstModule/internal/app/entity"
	"MyFirstModule/internal/app/infrastructure"
	"MyFirstModule/pkg"
)

type EventListener func(string)

type eventDispatcherUsecase interface {
	SetEventHandler(eventHandler EventListener)
	Start()
}

type fileReaderEventDispatcher struct {
	logger   pkg.Logger
	eventBus infrastructure.EventBus
	filePath string
}

func NewFileReaderEventDispatcher(logger pkg.Logger, eventBus infrastructure.EventBus, filePath string) *fileReaderEventDispatcher {
	return &fileReaderEventDispatcher{
		logger:   logger,
		eventBus: eventBus,
		filePath: filePath,
	}
}

func (fileReader *fileReaderEventDispatcher) SetEventHandler(eventHandler EventListener) {
	// fileReader.eventHandler = eventHandler
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
		fileReader.eventBus.Publish(entity.EventBusRawGameEvent, line)

		// fileReader.eventHandler(line)
	}
}
