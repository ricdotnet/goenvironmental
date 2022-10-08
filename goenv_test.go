package goenvironmental

import (
	"testing"
)

func TestParseEnv(t *testing.T) {
	ParseEnv(".env")

	if Get("HELLO") != "world" {
		t.Error("Failed simple string")
	}

	if Get("NUMBERS") != "0123456789" {
		t.Error("Failed numbers")
	}

	if Get("TEST5") != "hello=world" {
		t.Error("Failed more than one \"equals\"")
	}

	if Get("WITHQUOTES") != "hello world" {
		t.Error("Failed with quotes")
	}
}

func TestParseEnvDev(t *testing.T) {
	ParseEnv(".env.development")

	if Get("APP_SECRET") != "[this-is-some-app-secret]" {
		t.Error("Failed with custom file name")
	}
}
