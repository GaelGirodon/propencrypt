package test

import (
	"bytes"
	"gaelgirodon.fr/propencrypt/internal/app"
	"gaelgirodon.fr/propencrypt/internal/core/log"
	"os"
	"regexp"
	"testing"
)

func TestHelp(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantPattern string
		wantCode    int
	}{
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Mock output
			output := new(bytes.Buffer)
			log.Output = output
			// Set command-line arguments
			os.Args = []string{"propencrypt"}
			os.Args = append(os.Args, test.args...)
			// Run the application
			code := app.Run()
			// Assert
			if code != test.wantCode {
				t.Errorf("Expected code %d, got %d", test.wantCode, code)
			}
			regex := regexp.MustCompile(test.wantPattern)
			if !regex.Match(output.Bytes()) {
				t.Errorf("Expected output to match %s", test.wantPattern)
			}
		})
	}
}
