package singleton

import (
	"sync"
)

var lock = &sync.Mutex{}

type ApplicationContext struct {
    failedOperations int
}

var singleton *ApplicationContext

func getApplicationContext() *ApplicationContext {
    if singleton == nil {
        lock.Lock()
        defer lock.Unlock()

        if singleton == nil {
            singleton = &ApplicationContext{}
        }
    }

    return singleton
}

func (context *ApplicationContext) incrementFailedOperations() {
    context.failedOperations++
}

func (context *ApplicationContext) getFailedOperations() int {
    return context.failedOperations
}

func (context *ApplicationContext) reset() {
    context.failedOperations = 0
}