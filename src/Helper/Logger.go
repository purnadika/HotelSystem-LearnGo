package helper

import (
	"log"
	"time"
)

func LogDebug(message string, object interface{}) {
	printLog("DEBUG", message, object)
}
func LogWarning(message string, object interface{}) {
	printLog("WARNING", message, object)
}
func LogError(message string, object interface{}) {
	printLog("ERROR", message, object)
}

func printLog(logType string, message string, object interface{}) {
	log.Println(time.Now().String() + " [" + logType + "] " + message)
	objectString := SerializeObject(object)
	log.Println(objectString)
}
