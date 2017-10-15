package env

import (
	"os"
	"testing"
)

func TestGetEnvSetEnv(t *testing.T) {
	t.Log("Testing GetEnv with set environment variable T=test")
	os.Setenv("T", "test")
	a := GetEnv("T", "fail")
	switch a := interface{}(a).(type) {
	case string:
		if a != "test" {
			t.Errorf("GetEnv didn't return correct string from set environment variable")
		}
	default:
		t.Errorf("GetEnv didn't return a string\n")
	}
}

func TestGetEnvFallbackEnv(t *testing.T) {
	t.Log("Testing GetEnv with unset T: GetEnv(\"T\", \"fallback\")")
	os.Unsetenv("T")
	a := GetEnv("T", "fallback")
	switch a := interface{}(a).(type) {
	case string:
		if a != "fallback" {
			t.Errorf("GetEnv didn't return correct string from fallback")
		}
	default:
		t.Errorf("GetEnv didn't return a string\n")
	}
}
