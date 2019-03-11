package tomato

import (
	"fmt"
	"strings"
	"time"
)

type Timer struct {
	remaining time.Duration
}

func SetTimer(duration time.Duration) *Timer {
	return &Timer{remaining: duration}
}

func (this *Timer) Start() {
	fmt.Print(this)
	for this.remaining > 0 {
		time.Sleep(time.Second)
		this.remaining -= time.Second
		fmt.Print(this)
	}
	fmt.Print(clearLine)
}

func (this *Timer) String() string {
	return clearLine + this.remaining.String()
}

var clearLine = "\r" + strings.Repeat(" ", 8) + "\r"
