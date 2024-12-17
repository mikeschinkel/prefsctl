package cobrautil

import (
	"sync"
)

type Initializer func(cli *CLI)

var initializers = make([]Initializer, 0)
var initializerMutex sync.Mutex

func AddInitializer(init Initializer) {
	initializerMutex.Lock()
	initializers = append(initializers, init)
	initializerMutex.Unlock()
}
func RunInitializers(cli *CLI) {
	cobraRoot := RootCmd().Command()
	cobraRoot.ResetCommands()
	cobraRoot.ResetFlags()
	for _, initializer := range initializers {
		initializer(cli)
	}
}
