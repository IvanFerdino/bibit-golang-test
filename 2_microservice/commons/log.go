package commons

import (
	"log"
	"os"
)

func LogError(msg string) {
	log.SetOutput(os.Stderr)
	log.Printf(msg)
}

func LogInfo(msg string) {
	log.SetOutput(os.Stdout)
	log.Printf(msg)
}
