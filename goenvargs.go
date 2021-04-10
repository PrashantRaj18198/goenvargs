package goenvargs

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadEnvVars(
	expectedEnvVariables []string,
	splitter *string,
) (map[string]string, error) {

	defaultSplitter := "="
	if splitter == nil {
		splitter = &defaultSplitter
	}
	// the first arg is ./program ignore it
	argsWithoutProg := os.Args[1:]

	// create an empty map to store the parsed env variables
	envVariablesMap := make(map[string]string)

	// loop through all the arguments
	for _, arg := range argsWithoutProg {
		keyAndVal := strings.Split(arg, *splitter)

		// only parse arguments which has two elements
		if len(keyAndVal) == 2 {
			envVariablesMap[keyAndVal[0]] = keyAndVal[1]
			os.Setenv(keyAndVal[0], keyAndVal[1])
		}

	}

	// loop through all the expectedVariables
	for _, env := range expectedEnvVariables {
		// if the expected variable is not present return map with err
		if _, ok := envVariablesMap[env]; !ok {
			errString := fmt.Sprintf("Expected env variable %s was not found", env)
			return envVariablesMap, errors.New(errString)
		}
	}

	// return map and no error
	return envVariablesMap, nil
}
