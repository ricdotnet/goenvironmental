package goenvironmental

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
	rawEnv := string(envBytes)
	pairs := strings.Split(rawEnv, "\n")
	for _, pair := range pairs {

		// skip lines that does not follows key:value pair pattern
		if !strings.Contains(pair, "=") {
			continue
		}

		// strings.SplitN splits into a defined number of substrings
		// this allows us to have '=' in our values but not in keys
		parts := strings.SplitN(pair, "=", 2)
		var part = parts[1]

		// in case values have quotation marks we remove them
		if strings.HasPrefix(part, "\"") && strings.HasSuffix(part, "\"") {
			part = parts[1][1 : len(parts[1])-1]
		}

		os.Setenv(parts[0], part)
	}
}

func Get(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s does not exist", key)
	}

	return value, nil
}
