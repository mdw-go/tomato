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
	for range loop(config.TomatoesPerSet) {
		Tomato(config.WorkPeriod, config.RestPeriod, config.Silent)
		Prompt("Press <enter> to begin next tomato...")
	}
	Announce(FinishedTomato, config.Silent)
}

type Configuration struct {
	Silent         bool
	TomatoesPerSet int
	WorkPeriod     time.Duration
	RestPeriod     time.Duration
}

func ReadConfiguration() (config Configuration) {
	flag.BoolVar(&config.Silent, "silent", false, "When set, refrain from audio announcements.")
	flag.IntVar(&config.TomatoesPerSet, "tomatoes", 4, "How many tomatoes in this set?")
	flag.DurationVar(&config.WorkPeriod, "work", time.Minute*20, "How long is each work period?")
	flag.DurationVar(&config.RestPeriod, "rest", time.Minute*5, "How long is each rest period?")
	flag.Parse()
	return config
}

func loop(n int) []struct{} { return make([]struct{}, n) }

func Prompt(message string) {
	fmt.Print(message)
	bufio.NewScanner(os.Stdin).Scan()
}

func Tomato(work, rest time.Duration, silent bool) {
	Work(work)
	Rest(rest, silent)
}

func Work(work time.Duration) {
	fmt.Println(StartingTomato)
	tomato.SetTimer(work).Start()
}

func Rest(rest time.Duration, silent bool) {
	Announce(StoppingTomato, silent)
	MissionControl()
	tomato.SetTimer(rest).Start()
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

func Announce(message string, silent bool) {
	fmt.Println(message)
	if silent {
		Notify(message)
	} else {
		Execute("say", message)
	}
}

func Execute(command string, args ...string) {
	_ = exec.Command(command, args...).Start()
}

const (
	StartingTomato = "Starting tomato."
	StoppingTomato = "Tomato concluded; time for a break."
	FinishedTomato = "Tomato sessions complete; time for an extended break."
)
