package logging

import (
	"fmt"
	"os"
	"time"
)

func LogStd(message string) {
	Log(message, false, nil)
}

func LogError(message string, errMsg interface{}) {
	Log(message, true, errMsg)
}

func Log(message string, isError bool, err interface{}) {

	writer := os.Stdout
	var formattedMessage string

	if isError {
		writer = os.Stderr
		formattedMessage = fmt.Sprintf("[%s] Exception occurred! | Message: %s | Details: %v", time.Now().String(), message, err)
	} else {
		formattedMessage = fmt.Sprintf("[%s] %s", time.Now().String(), message)
	}

	fmt.Fprintln(writer, formattedMessage)
}
