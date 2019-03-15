package main

import (
	"flag"
	"time"

	"github.com/mdwhatcott/tomato"
	"github.com/mdwhatcott/tomato/external"
)

func main() {
	config := ReadConfiguration()
	session := NewSession(config)
	session.Run()
}

type Configuration struct {
	Silent         bool
	TomatoesPerSet int
	WorkPeriod     time.Duration
}

func ReadConfiguration() (config Configuration) {
	flag.BoolVar(&config.Silent, "silent", false, "When set, refrain from audio announcements.")
	flag.IntVar(&config.TomatoesPerSet, "tomatoes", 4, "How many tomatoes in this set?")
	flag.DurationVar(&config.WorkPeriod, "work", time.Minute*20, "How long is each work period?")
	flag.Parse()
	return config
}

type Session struct{ Configuration }

func NewSession(configuration Configuration) *Session {
	return &Session{Configuration: configuration}
}

func (this *Session) Run() {
	for x := range loop(this.TomatoesPerSet) {
		this.DoTomato(x == 0)
	}
	external.Announce(FinishedTomato, this.Silent)
}

func (this *Session) DoTomato(first bool) {
	if !first {
		external.Prompt(StartWhenReady)
	}
	Work(this.WorkPeriod, this.Silent)
	Rest(this.Silent)
}

func Work(work time.Duration, silent bool) {
	external.Announce(StartingTomato, silent)
	tomato.SetTimer(work).Start()
}

func Rest(silent bool) {
	external.Announce(StoppingTomato, silent)
	external.MissionControl()
}

const (
	StartingTomato = "Starting tomato."
	StoppingTomato = "Tomato concluded; time for a break."
	FinishedTomato = "Tomato sessions complete; time for an extended break."
	StartWhenReady = "Press <enter> to ready to begin the next tomato..."
)

func loop(n int) []struct{} { return make([]struct{}, n) }
