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
		WorkPeriodInMinutes     = flag.Duration("work", time.Minute * 25, "How long is each work period?")
		RestPeriodInMinutes     = flag.Duration("rest", time.Minute * 5, "How long is each rest period?")
		LongRestPeriodInMinutes = flag.Duration("long-rest", time.Minute * 15, "How long is the final rest period?")
	)

	flag.Parse()

	for x := 0; x < *TomatoesPerSet-1; x++ {
		Tomato(*WorkPeriodInMinutes, *RestPeriodInMinutes)
		fmt.Println("<ENTER> to begin the next tomato")
		fmt.Scanln()
	}
	Tomato(*WorkPeriodInMinutes, *LongRestPeriodInMinutes)
}

func Tomato(work, rest time.Duration) {
	Work(work)
	Rest(rest)
}

func Work(work time.Duration) {
	Notify(work.String() + " Tomato Starting")
	timer.SetTimer(work).Start()
}

func Rest(rest time.Duration) {
	Notify("Time for a " + rest.String() + " break")
	timer.SetTimer(rest).Start()
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
