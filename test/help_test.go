package test

import (
	"testing"
)

func TestHelp(t *testing.T) {
	tests := []commandTestCase{
		{name: "help/no-arg", args: []string{}, wantPattern: "propencrypt <command>", wantCode: 0},
		{name: "help/cmd-empty", args: []string{"help"}, wantPattern: "propencrypt <command>", wantCode: 0},
		{name: "help/opt", args: []string{"--help"}, wantPattern: "propencrypt <command>", wantCode: 0},
		{name: "help/opt-short", args: []string{"-h"}, wantPattern: "propencrypt <command>", wantCode: 0},
		{name: "help/cmd-encrypt", args: []string{"help", "encrypt"}, wantPattern: "propencrypt encrypt", wantCode: 0},
		{name: "help/cmd-decrypt", args: []string{"help", "decrypt"}, wantPattern: "propencrypt decrypt", wantCode: 0},
		{name: "help/cmd-unknown", args: []string{"help", "unknown"}, wantPattern: "unknown command", wantCode: 1},
		{name: "help/cmd-invalid", args: []string{"help", "a", "b"}, wantPattern: "invalid command", wantCode: 1},
		{name: "help/opt-encrypt", args: []string{"encrypt", "--help"}, wantPattern: "propencrypt encrypt", wantCode: 0},
		{name: "help/opt-decrypt", args: []string{"decrypt", "--help"}, wantPattern: "propencrypt decrypt", wantCode: 0},
		{name: "help/opt-encrypt-short", args: []string{"encrypt", "-h"}, wantPattern: "propencrypt encrypt", wantCode: 0},
		{name: "help/opt-decrypt-short", args: []string{"decrypt", "-h"}, wantPattern: "propencrypt decrypt", wantCode: 0},
		{name: "misc/unknown-cmd", args: []string{"unknown"}, wantPattern: "unknown command", wantCode: 1},
		{name: "misc/unknown-flag", args: []string{"--unknown"}, wantPattern: "unknown command", wantCode: 1},
		{name: "misc/unknown-cmd-flag", args: []string{"encrypt", "-?"}, wantPattern: "invalid command", wantCode: 1},
	}
	// Run command test cases
	testCommand(t, tests, []string{})
}
