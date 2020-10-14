package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mdwhatcott/tomato"
	"github.com/mdwhatcott/tomato/external"
)

func main() {
	NewSession(ReadConfiguration()).Run()
}

type Configuration struct {
	Silent         bool
	TomatoesPerSet int
	TeamSize       int
	WorkPeriod     time.Duration
}

func ReadConfiguration() (config Configuration) {
	flag.BoolVar(&config.Silent,
		"silent", false, "When set, refrain from audio announcements.",
	)
	flag.IntVar(&config.TomatoesPerSet,
		"tomatoes", 4, "How many tomatoes in this set?",
	)
	flag.IntVar(&config.TeamSize,
		"team", 1, "How many members in the team?",
	)
	flag.DurationVar(&config.WorkPeriod,
		"work", time.Minute*20, "How long is each work period?",
	)
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
			this.Rest(x)
		}
		this.Work(x + 1)
	}
	this.Finalize()
}

func (this *Session) Rest(session int) {
	external.Announce(fmt.Sprintf(StoppingTomato, session, this.TomatoesPerSet), this.Silent)
	this.prompt()
}

func (this *Session) prompt() {
	signal := make(chan bool)
	go this.awaitSignal(signal, time.Second*30)
	external.Announce(StartWhenReady, this.Silent)
	external.Prompt("")
	go this.send(signal)
}

func (this *Session) send(signal chan bool) {
	signal <- true
}

func (this *Session) awaitSignal(signal chan bool, delay time.Duration) {
	if delay > time.Minute * 5 {
		delay = time.Minute * 5
	}
	time.Sleep(delay)
	select {
	case <-signal:
		return
	default:
		external.MissionControl()
		external.Announce(StartWhenReady, this.Silent)
		this.awaitSignal(signal, delay+time.Minute)
	}
}
func (this *Session) Work(session int) {
	external.Announce(fmt.Sprintf(StartingTomato, session, this.TomatoesPerSet), this.Silent)
	for x := 0; x < this.TeamSize; x++ {
		tomato.SetTimer(this.WorkPeriod / time.Duration(this.TeamSize)).Start()
		if this.TeamSize > 1 && x < this.TeamSize-1 {
			external.Announce(SwitchDriver, this.Silent)
		}
	}
	external.MissionControl()
}
func (this *Session) Finalize() {
	external.Announce(fmt.Sprintf(FinishedTomato, this.TomatoesPerSet), this.Silent)
}

const (
	StartingTomato = "Starting tomato number %d of %d"
	StoppingTomato = "Tomato number %d of %d concluded; time for a break."
	FinishedTomato = "All %d tomato sessions complete; time for an extended break."
	StartWhenReady = "Press ENTER when ready to begin the next tomato..."
	SwitchDriver   = "Switch driver"
)
