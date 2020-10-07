package domain

import "log"

// FailOnError handle the error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Println("%s: %s", msg, err)
	}
}
