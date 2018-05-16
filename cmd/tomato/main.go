package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/mdwhatcott/tomato"
)

func main() {
	var (
		TomatoesPerSet          = flag.Int("tomatoes", 4, "How many tomatoes in this set?")
		WorkPeriodInMinutes     = flag.Int("work", 25, "How many minutes in each work period?")
		RestPeriodInMinutes     = flag.Int("rest", 5, "How many minutes in each rest period?")
		LongRestPeriodInMinutes = flag.Int("longrest", 15, "How many minutes in the final rest period?")
	)

	flag.Parse()

	for x := 0; x < *TomatoesPerSet-1; x++ {
		Tomato(*WorkPeriodInMinutes, *RestPeriodInMinutes)
		fmt.Println("<ENTER> to begin the next tomato")
		fmt.Scanln()
	}
	Tomato(*WorkPeriodInMinutes, *LongRestPeriodInMinutes)
}

func Tomato(work, rest int) {
	Work(work)
	Rest(rest)
}

func Work(minutes int) {
	workDuration := time.Minute * time.Duration(minutes)
	Notify(workDuration.String() + " Tomato Starting")
	timer.SetTimer(workDuration).Start()
}

func Rest(minutes int) {
	restDuration := time.Minute * time.Duration(minutes)
	Notify("Time for a " + restDuration.String() + " break")
	timer.SetTimer(restDuration).Start()
}

func Notify(message string) {
	fmt.Println(message)
	notification := fmt.Sprintf("display notification \"%s\" with title \"Tomato Timer\"", message)
	Execute(exec.Command("osascript", "-e", notification))
}

func Execute(command *exec.Cmd) {
	if output, err := command.CombinedOutput(); err != nil {
		log.Println("[WARN]", string(output), err)
	}
}
