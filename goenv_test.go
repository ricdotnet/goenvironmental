package goenvironmental

import (
	"testing"
)

func TestParseEnv(t *testing.T) {
	ParseEnv(".env")
	envs := EnvVariables

	if envs["HELLO"] != "world" {
		t.Error("Failed simple string")
	}

	if envs["NUMBERS"] != "0123456789" {
		t.Error("Failed numbers")
	}

	if envs["TEST5"] != "hello=world" {
		t.Error("Failed more than one \"equals\"")
	}

	if envs["WITHQUOTES"] != "hello world" {
		t.Error("Failed with quotes")
	}
}

func TestParseEnvDev(t *testing.T) {
	ParseEnv(".env.development")
	envs := EnvVariables

	if envs["APP_SECRET"] != "[this-is-some-app-secret]" {
		t.Error("Failed with custom file name")
	}
}
