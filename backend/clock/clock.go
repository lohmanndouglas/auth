package clock

import "time"

type Clock interface {
	CurrentMillis() int64
	Now() time.Time
}

type RealClock struct{}

func (rc *RealClock) CurrentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (rc *RealClock) Now() time.Time {
	return time.Now()
}
