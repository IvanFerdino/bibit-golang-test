package commons

import (
	"log"
	"os"
)

func LogError(msg string) {
	log.SetOutput(os.Stderr)
	log.Printf("[INFO] "+msg)
}

func LogInfo(msg string) {
	log.SetOutput(os.Stdout)
	log.Printf("[ERROR] "+msg)
}
