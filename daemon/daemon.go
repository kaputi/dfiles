package daemon

import "time"

type Daemon struct {
	running bool
	lastRun time.Time
}

var daemon = Daemon{
	running: false,
	lastRun: time.Time{},
}

func Status() {
	if daemon.running {
		println("Daemon is running")
		// TODO: rest of status
	} else {
		println("Daemon is stopped")
	}
}

func Start() {}

func Stop() {}
