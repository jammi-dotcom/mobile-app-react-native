package mobileappreactnative

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// GetBody reads the body of a request
func GetBody(r *http.Request) string {
	var body bytes.Buffer
	_, err := r.Body.ReadFrom(&body)
	if err != nil {
		log.Println(err)
	}
	return body.String()
}

// GetHeader reads the value of a header
func GetHeader(r *http.Request, header string) string {
	return r.Header.Get(header)
}

// GetQuery reads the value of a query parameter
func GetQuery(r *http.Request, query string) string {
	return r.URL.Query().Get(query)
}

// GetJSON reads the JSON body of a request
func GetJSON(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

// GetParams reads the parameters of a route
func GetParams(r *http.Request) map[string]string {
	vars := mux.Vars(r)
	return vars
}

// GetIntParam reads an integer parameter of a route
func GetIntParam(r *http.Request, param string) (int, error) {
	vars := GetParams(r)
	paramValue := vars[param]
	if paramValue == "" {
		return 0, fmt.Errorf("parameter '%s' was not provided", param)
	}
	intValue, err := strconv.Atoi(paramValue)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

// GetIntParamOrDefault reads an integer parameter of a route with a default value
func GetIntParamOrDefault(r *http.Request, param string, default int) int {
	intValue, err := GetIntParam(r, param)
	if err != nil {
		return default
	}
	return intValue
}

// GetBoolParam reads a boolean parameter of a route
func GetBoolParam(r *http.Request, param string) (bool, error) {
	vars := GetParams(r)
	paramValue := vars[param]
	if paramValue == "" {
		return false, fmt.Errorf("parameter '%s' was not provided", param)
	}
	if paramValue == "true" {
		return true, nil
	}
	if paramValue == "false" {
		return false, nil
	}
	return false, fmt.Errorf("invalid boolean value for parameter '%s'", param)
}

// GetBoolParamOrDefault reads a boolean parameter of a route with a default value
func GetBoolParamOrDefault(r *http.Request, param string, default bool) bool {
	boolValue, err := GetBoolParam(r, param)
	if err != nil {
		return default
	}
	return boolValue
}

// GetBoolOrDefault reads a boolean value from an environment variable with a default value
func GetBoolOrDefault(env string, default bool) bool {
	value := os.Getenv(env)
	if value == "" {
		return default
	}
	if value == "true" {
		return true
	}
	if value == "false" {
		return false
	}
	return false
}

// GetIntOrDefault reads an integer value from an environment variable with a default value
func GetIntOrDefault(env string, default int) int {
	value := os.Getenv(env)
	if value == "" {
		return default
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return default
	}
	return intValue
}

// GetFloatOrDefault reads a float value from an environment variable with a default value
func GetFloatOrDefault(env string, default float64) float64 {
	value := os.Getenv(env)
	if value == "" {
		return default
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return default
	}
	return floatValue
}

// GetStringsOrDefault reads a list of strings from an environment variable with a default value
func GetStringsOrDefault(env string, default []string) []string {
	value := os.Getenv(env)
	if value == "" {
		return default
	}
	stringsSlice := strings.Split(value, ",")
	for i, s := range stringsSlice {
		stringsSlice[i] = strings.TrimSpace(s)
	}
	return stringsSlice
}

// GetBoolOr returns the first truthy value from a list
func GetBoolOr(values ...bool) bool {
	for _, value := range values {
		if value {
			return value
		}
	}
	return false
}

// GetAnyOr returns the first non-empty value from a list
func GetAnyOr(values ...any) any {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}