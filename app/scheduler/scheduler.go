package scheduler

import "log"

/*
Start func() error
*/
func Start() error {
	log.Println("Monitor Start...")
	err := startCron()
	return err
}

/*
Stop func() error
*/
func Stop() {
	log.Println("Monitor Stop...")
	C.Stop()
}

// func Clear() {
// 	C.
// }
