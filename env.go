package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type EnvironmentVariables struct {
	FilePath string
	Vars     map[string]string
}

func (vars *EnvironmentVariables) Load() {
	vars.Vars = make(map[string]string)

	body, err := os.ReadFile(vars.FilePath)
	if err != nil {
		panic(err)
	}

	str := string(body)
	lines := strings.Split(str, "\n")
	var keyValue []string
	var key, value string

	for i := 0; i < len(lines); i++ {
		keyValue = strings.Split(lines[i], "=")
		key = strings.ToLower(keyValue[0])
		value = keyValue[1]

		vars.Vars[key] = value
	}
}

func (vars *EnvironmentVariables) GetString(key string) string {
	key = strings.ToLower(key)
	value := vars.Vars[key]

	if value == "" {
		log.Printf("WARN::ENV Unknown Environment Variable \"%s\"", key)
		return ""
	}

	return vars.Vars[key]
}

func (vars *EnvironmentVariables) GetBoolean(key string) bool {
	key = strings.ToLower(key)
	value := vars.Vars[key]

	if value == "" {
		log.Printf("WARN::ENV Unknown Environment Variable \"%s\"", key)
		return false
	}

	return vars.Vars[key][0] == 't' || vars.Vars[key][0] == 'T'
}

func (vars *EnvironmentVariables) GetInteger(key string) int {
	key = strings.ToLower(key)

	val, err := strconv.Atoi(vars.Vars[key])
	if err != nil {
		log.Printf("WARN::ENV Attempted to parse an environment variable (%s=%s) to an integer unsuccessfully", key, vars.Vars[key])
		return -1
	}

	return val
}
