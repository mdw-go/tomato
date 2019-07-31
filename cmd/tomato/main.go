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
	TeamSize       int
	WorkPeriod     time.Duration
}

func ReadConfiguration() (config Configuration) {
	flag.BoolVar(&config.Silent, "silent", false, "When set, refrain from audio announcements.")
	flag.IntVar(&config.TomatoesPerSet, "tomatoes", 4, "How many tomatoes in this set?")
	flag.IntVar(&config.TeamSize, "team", 1, "How many members in the team?")
	flag.DurationVar(&config.WorkPeriod, "work", time.Minute*20, "How long is each work period?")
	flag.Parse()
	return config
}

type Session struct{ Configuration }

func NewSession(configuration Configuration) *Session {
	return &Session{Configuration: configuration}
}

func (this *Session) Run() {
	for x := 0; x < this.TomatoesPerSet; x++ {
		if x > 0 {
			this.Rest()
		}
		this.Work()
	}
	this.Finalize()
}

func (this *Session) Rest() {
	external.Announce(StoppingTomato, this.Silent)
	external.Prompt(StartWhenReady)
}
func (this *Session) Work() {
	external.Announce(StartingTomato, this.Silent)
	for x := 0; x < this.TeamSize; x++ {
		tomato.SetTimer(this.WorkPeriod / time.Duration(this.TeamSize)).Start()
		if this.TeamSize > 1 && x < this.TeamSize-1 {
			external.Announce(SwitchDriver, this.Silent)
		}
	}
	external.MissionControl()
}
func (this *Session) Finalize() {
	external.Announce(FinishedTomato, this.Silent)
}

const (
	StartingTomato = "Starting tomato."
	StoppingTomato = "Tomato concluded; time for a break."
	FinishedTomato = "Tomato sessions complete; time for an extended break."
	StartWhenReady = "Press <enter> to ready to begin the next tomato..."
	SwitchDriver   = "Switch driver"
)
