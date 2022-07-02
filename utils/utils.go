package utils

import (
	"log"
	"math/rand"
)

func FailIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateUid() int64 {
	return rand.Int63()
}
