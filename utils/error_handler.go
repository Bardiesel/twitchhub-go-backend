package utils

import "log"


func ErrorHandler(err error, msg string) {
	if err != nil {
		log.Println(msg)
	}
}