package sleeper

import (
	"time"
)

type Sleeper struct {
	duration time.Duration
}

func New(d time.Duration) *Sleeper {
	return &Sleeper{duration: d}
}

func (s *Sleeper) Sleep() {
	time.Sleep(s.duration * time.Second)
}
