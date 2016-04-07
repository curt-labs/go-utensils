package speak

import (
	"errors"
	"os/exec"
)

var Voices = []string{"Agnes", "Kathy", "Princess",
	"Vicki", "Victoria", "Bruce",
	"Fred", "Junior", "Ralph",
	"Albert", "Bahh", "Bells",
	"Boing", "Bubbles", "Cellos",
	"Deranged", "Hysterical", "Trinoids",
	"Whisper", "Zarvox"}

//Speak Spits out string message to console to hear.
func Speak(str, voice string) error {
	if str == "" {
		return errors.New("Problem executing: No message provided.")
	}
	cmd := exec.Command("say", "-v", GetVoice(voice), str)
	return cmd.Run()
}

func GetVoice(req string) string {
	for _, v := range Voices {
		if req == v {
			return v
		}
	}
	return "Alex"
}
