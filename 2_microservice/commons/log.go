package commons

import (
	"log"
	"os"
)

func LogError(msg string) {
	log.SetOutput(os.Stderr)
	log.Printf("[ERROR] "+msg)
}

func LogInfo(msg string) {
	log.SetOutput(os.Stdout)
	log.Printf("[INFO] "+msg)
}
