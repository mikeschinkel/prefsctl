package cobrautil

import (
	"sync"
)

var result Result
var resultMutex sync.Mutex

func SetResult(r Result) {
	resultMutex.Lock()
	result = r
	resultMutex.Unlock()
}
