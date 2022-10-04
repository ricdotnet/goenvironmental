package goenvironmental

import (
	"testing"
)

func TestParseEnv(t *testing.T) {
	ParseEnv()
	envs := EnvVariables

	if envs["HELLO"] != "world" {
		t.Error("Invalid string value")
	}

	if envs["NUMBERS"] != "0123456789" {
		t.Error("Invalid numbers value")
	}

	if envs["TEST5"] != "hello=world" {
		t.Error("Invalid value")
	}
}
