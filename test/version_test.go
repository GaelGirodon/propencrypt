package test

import (
	"bytes"
	"gaelgirodon.fr/propencrypt/internal/app"
	"gaelgirodon.fr/propencrypt/internal/core/log"
	"os"
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantPattern string
		wantCode    int
	}{
		{name: "version/cmd", args: []string{"version"}, wantPattern: "propencrypt version [0-9.]+", wantCode: 0},
		{name: "version/opt", args: []string{"--version"}, wantPattern: "propencrypt version [0-9.]+", wantCode: 0},
		{name: "version/opt-short", args: []string{"-v"}, wantPattern: "propencrypt version [0-9.]+", wantCode: 0},
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
