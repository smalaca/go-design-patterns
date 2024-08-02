package singleton

type FailuresStorage struct {
    failedOperations int
}

func (storage *FailuresStorage) incrementFailedOperations() {
    storage.failedOperations++
}

func (storage *FailuresStorage) getFailedOperations() int {
    return storage.failedOperations
}

func (storage *FailuresStorage) reset() {
    storage.failedOperations = 0
}