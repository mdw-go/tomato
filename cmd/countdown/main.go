package main

import (
	"flag"
	"time"

	"github.com/mdw-go/tomato"
	"github.com/mdw-go/tomato/external"
)

func main() {
	var duration time.Duration
	var silent bool
	flag.DurationVar(&duration, "duration", time.Minute, "The initial value of the countdown timer.")
	flag.BoolVar(&silent, "silent", false, "When set, silence final announcement.")
	flag.Parse()
	tomato.SetTimer(duration).Start()
	external.Announce("Countdown complete.", silent)
}
