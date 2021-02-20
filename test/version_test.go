package test

import (
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []commandTestCase{
		{name: "version/cmd", args: []string{"version"}, wantPattern: "propencrypt version [0-9.]+", wantCode: 0},
		{name: "version/opt", args: []string{"--version"}, wantPattern: "propencrypt version [0-9.]+", wantCode: 0},
		{name: "version/opt-short", args: []string{"-v"}, wantPattern: "propencrypt version [0-9.]+", wantCode: 0},
	}
	// Run command test cases
	testCommand(t, tests, []string{})
}
