package utils

type Mutex interface {
	Lock()
	TryLock() bool
	Unlock()
}
