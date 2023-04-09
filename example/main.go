package main

import (
	"fmt"
	goenv "github.com/ricdotnet/goenvironmental"
	"github.com/ricdotnet/goenvironmental/example/more"
)

func main() {
	goenv.ParseEnv()
	println(goenv.Get("HELLO"))
	println(goenv.Get("THIS"))
	println(goenv.Get("NUMBERS"))
	println(goenv.Get("TEST5"))
	println(goenv.Get("WITHQUOTES"))
	println(goenv.Get("NUMBERS"))

	more.Hello()

	// if an env var does not exist
	_, err := goenv.Get("SDKJSDFSDG")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
