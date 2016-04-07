package speak

import (
	"log"
	"os/exec"
)

//Speak Spits out string message to console to hear.
func Speak(str string) {
	if str == "" {
		log.Println("Problem executing: No message provided.")
	}
	cmd := exec.Command("say", str)
	err := cmd.Run()

	if err != nil {
		log.Printf("Problem executing: %s", err)
	}
}
