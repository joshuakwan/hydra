package utils

import (
	"log"
	"os/exec"
)

// GenerateUUID generates a unique UUID
func GenerateUUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
