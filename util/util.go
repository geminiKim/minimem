package util

import (
	"strconv"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetInt(value string) int {
	intValue, _ := strconv.Atoi(value)
	return intValue
}

func ToJsonString(value []string) string {
	json, _ := json.Marshal(value)
	return string(json)
}

func GetMessageMapWithBody(request *http.Request, uriKeys []string, bodyKey string) map[string]string {
	value := GetMessageMap(request,uriKeys)
	value[bodyKey] = readBody(request)
	return value
}

func GetMessageMap(request *http.Request, uriKeys []string) map[string]string {
	value := make(map[string]string)

	vars := mux.Vars(request)
	for _, key := range uriKeys {
		value[key] = vars[key]
	}
	return value
}

func readBody(request *http.Request) string {
	result, _ := ioutil.ReadAll(request.Body)
	return string(result)
}
