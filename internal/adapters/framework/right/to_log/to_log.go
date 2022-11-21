package to_log

import (
	"log"
	"time"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (pa Adapter) CloseDbConnection() {}

func (pa Adapter) AddToHistory(answer int32, operation string) error {
	log.Printf("Got %v result: %v at %v", operation, answer, time.Now())
	return nil
}
