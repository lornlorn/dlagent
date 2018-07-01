package api

import "log"

// API struct
type API struct{}

// GetJobStatus (api API) func(name string) error
func (api API) GetJobStatus(a string, b string) error {
	log.Println(a, b)
	return nil
}
