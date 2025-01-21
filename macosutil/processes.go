package macosutil

type processesMap map[DomainName]ProcessName

var processesToKill = processesMap{
	"com.apple.dock": "Dock",
}

func GetProcessToKill(name DomainName) ProcessName {
	process, _ := processesToKill[name]
	return process
}
