package helper

import (
	log "github.com/sirupsen/logrus"
)

// Here we set the way error messages are displayed in the terminal.
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
