package helper

import (
	"encoding/json"
	"strconv"
)

func StringToUint(uintString string) uint {
	result, err := strconv.ParseUint(uintString, 10, 32)
	PanicIfError(err)
	return uint(result)
}
func UintToString(value uint) string {
	result := strconv.FormatUint(uint64(value), 10)
	return result
}

func SerializeObject(object interface{}) string {
	result, err := json.Marshal(object)
	PanicIfError(err)
	return string(result)
}
