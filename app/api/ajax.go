package api

import (
	"app/scheduler"
	"log"
)

// API struct
type API struct {
}

// StopScheduler (api API) func(a []byte) ([]byte, error)
func (api API) StopScheduler(a []byte) ([]byte, error) {
	log.Println(a)

	scheduler.Stop()

	return []byte("StopScheduler"), nil
}
