package internal

type Mutex interface {
	Lock()
	TryLock() bool
	Unlock()
}
