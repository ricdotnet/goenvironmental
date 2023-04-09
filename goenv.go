package goenvironmental

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var envVariables map[string]string

// ParseEnv reads the contents from .env or other defined file
// place the .env file on the root folder
func ParseEnv(args ...string) {

	if len(args) > 1 {
		log.Fatal("ParseEnv accepts one and only one file name")
		return
	}

	var fileName = ".env"
	if len(args) == 1 {
		fileName = args[0]
	}

	// envBytes will be the content of the .env file in bytes
	envBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	// generate a string from envBytes
	// then split the content by \n into lines
	// after we have to get key:value pairs and store in a map for later use
	envVariables = make(map[string]string)
	rawEnv := string(envBytes)
	pairs := strings.Split(rawEnv, "\n")
	for _, pair := range pairs {

		// strings.SplitN splits into a defined number of substrings
		// this allows us to have '=' in our values but not in keys
		parts := strings.SplitN(pair, "=", 2)
		var part = parts[1]

		// in case values have quotation marks we remove them
		if strings.HasPrefix(part, "\"") && strings.HasSuffix(part, "\"") {
			part = parts[1][1 : len(parts[1])-1]
		}

		envVariables[strings.TrimRight(parts[0], "=")] = part
	}
}

func Get(key string) (string, error) {
	var value string
	value = envVariables[key]

	if value == "" {
		return "", fmt.Errorf("%s does not exist", key)
	}

	return value, nil
}
