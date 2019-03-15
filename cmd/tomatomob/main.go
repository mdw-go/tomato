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
	TeamSize       int
}

func ReadConfiguration() (config Configuration) {
	flag.BoolVar(&config.Silent, "silent", false, "When set, refrain from audio announcements.")
	flag.IntVar(&config.TomatoesPerSet, "tomatoes", 4, "How many tomatoes in this set?")
	flag.IntVar(&config.TeamSize, "team", 3, "How many members in the team?")
	flag.DurationVar(&config.WorkPeriod, "work", time.Minute*20, "How long is each work period?")
	flag.Parse()
	return config
}

type Session struct {
	Configuration
}

func NewSession(configuration Configuration) *Session {
	return &Session{Configuration: configuration}
}

func (this *Session) Run() {
	for x := range loop(this.TomatoesPerSet) {
		this.DoTomato(x == 0, x == this.TomatoesPerSet-1)
	}
}

func (this *Session) DoTomato(first bool, last bool) {
	if !first {
		external.Prompt(StartWhenReady)
	}

	this.Work()

	if !last {
		Rest(this.Silent)
	} else {
		external.Announce(FinishedTomato, this.Silent)
	}
}

func (this *Session) Work() {
	external.Announce(StartingTomato, this.Silent)
	for x := range loop(this.TeamSize) {
		tomato.SetTimer(this.WorkPeriod / time.Duration(this.TeamSize)).Start()
		if x < this.TeamSize-1 {
			external.Announce(SwitchDriver, this.Silent)
		}
	}
}

func Rest(silent bool) {
	external.Announce(StoppingTomato, silent)
	external.MissionControl()
}

const (
	StartingTomato = "Starting tomato."
	StoppingTomato = "Tomato concluded; time for a break."
	FinishedTomato = "Tomato sessions complete; time for an extended break."
	SwitchDriver   = "Switch driver"
	StartWhenReady = "Press <enter> to ready to begin the next tomato..."
)

func loop(n int) []struct{} { return make([]struct{}, n) }
