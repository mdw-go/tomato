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
	for range loop(config.TomatoesPerSet - 1) {
		Tomato(config.WorkPeriod, config.RestPeriod)
		Prompt("Press <enter> to begin next tomato...")
	}
	Tomato(config.WorkPeriod, config.LongRestPeriod)
}

type Configuration struct {
	TomatoesPerSet int
	WorkPeriod     time.Duration
	RestPeriod     time.Duration
	LongRestPeriod time.Duration
}

func ReadConfiguration() (config Configuration) {
	flag.IntVar(&config.TomatoesPerSet, "tomatoes", 4, "How many tomatoes in this set?")
	flag.DurationVar(&config.WorkPeriod, "work", time.Minute*20, "How long is each work period?")
	flag.DurationVar(&config.RestPeriod, "rest", time.Minute*5, "How long is each rest period?")
	flag.DurationVar(&config.LongRestPeriod, "long-rest", time.Minute*10, "How long is the final rest period?")
	flag.Parse()
	return config
}

func loop(n int) []struct{} { return make([]struct{}, n) }

func Prompt(message string) {
	fmt.Print(message)
	bufio.NewScanner(os.Stdin).Scan()
}

func Tomato(work, rest time.Duration) {
	Work(work)
	Rest(rest)
}

func Work(work time.Duration) {
	Notify(work.String() + " Tomato Starting")
	tomato.SetTimer(work).Start()
}

func Rest(rest time.Duration) {
	Distract()
	Notify("Time for a " + rest.String() + " break")
	tomato.SetTimer(rest).Start()
}

func Notify(message string) {
	fmt.Println(message)
	AppleScript(fmt.Sprintf("display notification \"%s\" with title \"Tomato Timer\"", message))
}

func Distract() {
	AppleScript("tell application \"Mission Control\" to activate")
}

func AppleScript(script string) {
	exec.Command("osascript", "-e", script).Start()
}
