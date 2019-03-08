package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/mdwhatcott/tomato"
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

type Session struct {
	Configuration
}

func NewSession(configuration Configuration) *Session {
	return &Session{Configuration: configuration}
}

func (this *Session) Run() {
	for x := range loop(this.TomatoesPerSet) {
		this.DoTomato(x == 0)
	}
	Announce(FinishedTomato, this.Silent)
}

func (this *Session) DoTomato(first bool) {
	if !first {
		Prompt("Press <enter> to ready to begin the next tomato...")
	}
	Work(this.WorkPeriod, this.Silent)
	Rest(this.Silent)
}

func Tomato(work time.Duration, silent bool) {
	Work(work, silent)
	Rest(silent)
}

func Work(work time.Duration, silent bool) {
	Announce(StartingTomato, silent)
	tomato.SetTimer(work).Start()
}

func Rest(silent bool) {
	Announce(StoppingTomato, silent)
	MissionControl()
}

func Announce(message string, silent bool) {
	fmt.Println(message)
	if silent {
		Notify(message)
	} else {
		Execute("say", message)
	}
}

func Notify(message string) {
	AppleScript(fmt.Sprintf("display notification \"%s\" with title \"Tomato Timer\"", message))
}

func MissionControl() {
	AppleScript("tell application \"Mission Control\" to activate")
}

func AppleScript(script string) {
	Execute("osascript", "-e", script)
}

func Execute(command string, args ...string) {
	_ = exec.Command(command, args...).Start()
}

func Prompt(message string) {
	fmt.Print(message)
	bufio.NewScanner(os.Stdin).Scan()
}

const (
	StartingTomato = "Starting tomato."
	StoppingTomato = "Tomato concluded; time for a break."
	FinishedTomato = "Tomato sessions complete; time for an extended break."
)

func loop(n int) []struct{} { return make([]struct{}, n) }
