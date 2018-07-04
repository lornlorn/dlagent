package api

import (
	"app/scheduler"
	"log"
)

// API struct
type API struct {
}

// GetJobStatus (api API) func(name string) error
func (api API) StopScheduler(a []byte) error {
	log.Println(a)

	scheduler.Stop()

	return nil
}
