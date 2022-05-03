package utils

import "log"

func failIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
