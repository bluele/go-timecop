package timecop

import (
	"sync"
	"time"
)

var (
	isFreeze   = false
	isTravel   = false
	freezeTime time.Time
	travelTime time.Time

	useLock = true
	rw      sync.RWMutex
)

func UseLock() {
	rw.Lock()
	defer rw.Unlock()
	useLock = true
}

func UnUseLock() {
	rw.Lock()
	defer rw.Unlock()
	useLock = false
}

func Now() time.Time {
	if isFreeze {
		if useLock {
			rw.RLock()
			defer rw.RUnlock()
		}
		return freezeTime
	}

	if isTravel {
		if useLock {
			rw.RLock()
			defer rw.RUnlock()
		}
		return freezeTime.Add(time.Now().Sub(travelTime))
	}

	return time.Now()
}

func Freeze(t time.Time) {
	if useLock {
		rw.Lock()
		defer rw.Unlock()
	}
	freezeTime = t
	isFreeze = true
}

func Travel(t time.Time) {
	if useLock {
		rw.Lock()
		defer rw.Unlock()
	}
	freezeTime = t
	travelTime = time.Now()
	isTravel = true
}

func Since(t time.Time) time.Duration {
	return Now().Sub(t)
}

func Return() {
	if useLock {
		rw.Lock()
		defer rw.Unlock()
	}
	isFreeze = false
	isTravel = false
}
