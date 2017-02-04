package util

import (
	"strconv"
	"io/ioutil"
	"net/http"
)

func GetInt(value string) int {
	intValue, _ := strconv.Atoi(value)
	return intValue
}

func ReadAll(request *http.Request) string {
	result, _ := ioutil.ReadAll(request.Body)
	return string(result)
}


