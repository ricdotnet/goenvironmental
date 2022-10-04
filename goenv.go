package goenvironmental

import (
	"io/ioutil"
	"log"
	"strings"
)

var EnvVariables map[string]string

// ParseEnv reads the contents from .env file
// place the .env file on the root folder
func ParseEnv() {

	// c will be the content of the .env file in bytes
	c, err := ioutil.ReadFile(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	// generate content from c
	// then split the content by line
	// after we have to get key:value pairs and store in a map for later use
	EnvVariables = make(map[string]string)
	content := string(c)
	pairs := strings.Split(content, "\n")
	for _, pair := range pairs {

		// strings.SplitN splits into a defined number of substrings
		// this allows us to have '=' in our values but not in keys
		kV := strings.SplitN(pair, "=", 2)
		EnvVariables[strings.TrimRight(kV[0], "=")] = kV[1]
	}
}
