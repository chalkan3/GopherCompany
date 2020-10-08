package domain

import (
	"encoding/json"
	"log"
)

// FailOnError handle the error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Println("%s: %s", msg, err)
	}
}

func JSONStringfy(i interface{}) string {

	e, err := json.Marshal(i)
	FailOnError(err, "JSONStringfy")
	return string(e)

}
