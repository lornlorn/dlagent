package monitor

/*
Start func() error
*/
func Start() error {
	err := startCron()
	return err
}

/*
Stop func() error
*/
func Stop() {
	C.Stop()
}

// func Clear() {
// 	C.
// }
