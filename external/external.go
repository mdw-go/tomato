package external

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Announce(message string, silent bool) {
	fmt.Println(message)
	if silent {
		Notify(message)
	} else {
		Say(message)
	}
}

func Notify(message string) {
	AppleScript(fmt.Sprintf(`display notification "%s" with title "Tomato Timer"`, message))
}

func Say(message string) {
	Execute("say", "-v", "Samantha", message)
}

func MissionControl() {
	AppleScript(`tell application "Mission Control" to activate`)
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
