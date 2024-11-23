package cliutil

import (
	"sync"
)

type initializer func(cli *CLI)

var initializers = make([]initializer, 0)
var initializerMutex sync.Mutex

func AddInitializer(init initializer) {
	initializerMutex.Lock()
	initializers = append(initializers, init)
	initializerMutex.Unlock()
}
func RunInitializers(cli *CLI) {
	RootCmd().ResetCommands()
	RootCmd().ResetFlags()
	for _, initializer := range initializers {
		initializer(cli)
	}
}
