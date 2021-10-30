package main

import (
	goenv "github.com/ricdotnet/goenvironmental"
)

func main() {
	goenv.ParseEnv()
	envs := goenv.EnvVariables
	println(envs["HELLO"])
	println(envs["THIS"])
	println(envs["NUMBERS"])
	println(envs["TEST5"])
}
