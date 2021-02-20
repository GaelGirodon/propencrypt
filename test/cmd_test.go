package test

import (
	"bytes"
	"gaelgirodon.fr/propencrypt/internal/app"
	"gaelgirodon.fr/propencrypt/internal/core/log"
	"gaelgirodon.fr/propencrypt/pkg/fileutil"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

// commandTestCase is a test case for a command.
type commandTestCase struct {
	name            string
	args            []string
	env             map[string]string
	wantPattern     string
	wantCode        int
	wantFile        string
	wantFilePattern string
}

// testCommand runs command test cases.
func testCommand(t *testing.T, tests []commandTestCase, dataFiles []string) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Mock output
			output := new(bytes.Buffer)
			log.Output = output
			// Set environment variables
			for key, value := range test.env {
				_ = os.Setenv(key, value)
			}
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
				t.Errorf("Expected output to match '%s', got '%s'", test.wantPattern, output.Bytes())
			}
			if test.wantCode == 0 && len(test.wantFile) > 0 {
				if !fileutil.ExistsRegular(test.wantFile) {
					t.Errorf("Expected '%s' to exist", test.wantFile)
				}
				content, _ := ioutil.ReadFile(test.wantFile)
				if !regexp.MustCompile(test.wantFilePattern).Match(content) {
					t.Errorf("Expected '%s' to be valid", test.wantFile)
				}
			}
			// Remove output test data files
			for _, f := range dataFiles {
				if fileutil.ExistsRegular("data/" + f) {
					_ = os.Remove("data/" + f)
				}
			}
			// Clean up environment variables
			for key := range test.env {
				_ = os.Unsetenv(key)
			}
		})
	}
}
