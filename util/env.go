package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Environment struct
type Environment struct {
	config map[string]interface{}
}

// Set environment variable
func (e *Environment) Set(key, value string) {
	e.config[key] = value
	os.Setenv(key, value)
}

// Get environment variable
func (e *Environment) Get(key string) string {
	return os.Getenv(key)
}

// Load environment from file
func (e *Environment) Load() {
	e.loadFile(false)
}

// Override load environment and replace existing
func (e *Environment) Override() {
	e.loadFile(true)
}

// GetConfig return desired config
func (e *Environment) GetConfig() map[string]interface{} {
	return e.config
}

func (e *Environment) loadFile(overload bool) {
	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range e.config {
		if !currentEnv[key] || overload {
			val := fmt.Sprintf("%v", value)
			os.Setenv(key, val)
		}
	}
}

// Env create new Environment
func Env(path string) Environment {
	var environment Environment
	environment = Environment{Config(path)}
	environment.Load()

	return environment
}

// Config set application config
func Config(filepath string) map[string]interface{} {
	configs := make(map[string]interface{})
	pathWithFilename := filepath + ".env"
	b, err := ioutil.ReadFile(pathWithFilename)
	if err == nil {
		str := string(b)

		ArrayLines := strings.Split(str, "\n")
		for _, val := range ArrayLines {
			val = strings.TrimSpace(val)
			if val != "" && !strings.HasPrefix(val, "#") {
				explode := strings.Split(val, "=")
				configs[explode[0]] = explode[1]
				os.Setenv(explode[0], fmt.Sprint(explode[1]))
			}
		}
	}

	return configs
}
