package goenvironmental

import (
	"os"
	"testing"
)

func TestParseEnv(t *testing.T) {
	ParseEnv(".env")

	if s, _ := Get("HELLO"); s != "world" {
		t.Error("Failed simple string")
	}

	if s, _ := Get("NUMBERS"); s != "0123456789" {
		t.Error("Failed numbers")
	}

	if s, _ := Get("TEST5"); s != "hello=world" {
		t.Error("Failed more than one \"equals\"")
	}

	if s, _ := Get("WITHQUOTES"); s != "hello world" {
		t.Error("Failed with quotes")
	}

	if _, e := Get("NON_EXISTENT"); e == nil {
		t.Error("Failed non existent")
	}
}

func TestParseEnvDev(t *testing.T) {
	ParseEnv(".env.development")

	if s, _ := Get("APP_SECRET"); s != "[this-is-some-app-secret]" {
		t.Error("Failed with custom file name")
	}
}

func TestNoFile(t *testing.T) {
	ParseEnv()

	os.Setenv("EXISTING_VAR", "hello_world")

	if _, e := Get("EXISTING_VAR"); e != nil {
		t.Error("Failed existing var")
	}

	if _, e := Get("NO_ENV_AVAILABLE"); e == nil {
		t.Error("Failed no env available")
	}
}
