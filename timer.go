package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	seconds int
	tick    time.Duration
}

func SetTimer(duration time.Duration) *Timer {
	return &Timer{
		seconds: int(duration.Seconds()),
		tick:    time.Second,
	}
}

func (this *Timer) Start() {
	fmt.Print(this)
	for this.seconds > 0 {
		time.Sleep(this.tick)
		this.seconds--
		fmt.Print(this)
	}
	fmt.Println()
}

func (this *Timer) String() string {
	seconds := this.seconds

	minutes := seconds / 60
	seconds -= minutes * 60

	hours := minutes / 60
	minutes -= hours * 60

	if hours == 0 {
		if minutes == 0 {
			return fmt.Sprintf("\r%02d", seconds)
		}
		return fmt.Sprintf("\r%02d:%02d", minutes, seconds)
	}
	return fmt.Sprintf("\r%02d:%02d:%02d", hours, minutes, seconds)
}
