package macosutil

func KillDock() (err error) {
	return KillProcess("Dock")
}
