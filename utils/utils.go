package utils

import "log"

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %v", msg, err)
	}
}
