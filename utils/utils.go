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

func GenerateRandomNumber(max float64) float64 {
	return rand.Float64() * max
}
