package main

import (
	goenv "github.com/ricdotnet/goenvironmental"
)

func main() {
	goenv.ParseEnv()
	println(goenv.EnvVariables["HELLO"])
	println(goenv.EnvVariables["THIS"])
	println(goenv.EnvVariables["NUMBERS"])
	println(goenv.EnvVariables["TEST5"])
}
