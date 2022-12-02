package utils

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %v", msg, err)
	}
}

func HashingPassword(str string) string {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	return hash
}
