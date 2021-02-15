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

const (
	// key is a sample encryption key.
	key = "e2efc125f1cd4df089819271b75b81d5"

	// pattern is a sample pattern.
	pattern = "pass: (.+)"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantPattern string
		wantCode    int
	}{
		{name: "encrypt/opt-key-missing", args: []string{"encrypt"},
			wantPattern: "key is required", wantCode: 1},
		{name: "encrypt/opt-key-invalid", args: []string{"encrypt", "-k", "invalid"},
			wantPattern: "key is required", wantCode: 1},
		{name: "encrypt/opt-pattern-missing", args: []string{"encrypt", "-k", key},
			wantPattern: "pattern is required", wantCode: 1},
		{name: "encrypt/opt-pattern-no-group", args: []string{"encrypt", "-k", key, "-p", "invalid"},
			wantPattern: "no capturing group", wantCode: 1},
		{name: "encrypt/opt-pattern-invalid", args: []string{"encrypt", "-k", key, "-p", "[p=(.+)"},
			wantPattern: "pattern is invalid", wantCode: 1},
		{name: "encrypt/opt-ext-invalid", args: []string{"encrypt", "-k", key, "-p", pattern, "-e", "invalid"},
			wantPattern: "extension is invalid", wantCode: 1},
		{name: "encrypt/args-files-missing", args: []string{"encrypt", "-k", key, "-p", pattern},
			wantPattern: "missing files", wantCode: 1},
		{name: "encrypt/args-files-expand-nok", args: []string{"encrypt", "-k", key, "-p", pattern, "data/["},
			wantPattern: "invalid files", wantCode: 1},
		{name: "encrypt/args-files-expand-missing", args: []string{"encrypt", "-k", key, "-p", pattern, "nowhere/*.ini"},
			wantPattern: "missing files", wantCode: 1},
		{name: "encrypt/args-files-not-regular", args: []string{"encrypt", "-k", key, "-p", pattern, "data"},
			wantPattern: "not a regular file", wantCode: 1},
		{name: "encrypt/ok", args: []string{"encrypt", "-k", key, "-p", pattern, "data/enc.yml"},
			wantPattern: "", wantCode: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Mock output
			output := new(bytes.Buffer)
			log.Output = output
			// Set command-line arguments
			os.Args = []string{"propencrypt"}
			os.Args = append(os.Args, test.args...)
			// Remove output file
			if fileutil.ExistsRegular("data/enc.yml.enc") {
				_ = os.Remove("data/enc.yml.enc")
			}
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
			if test.wantCode == 0 {
				if !fileutil.ExistsRegular("data/enc.yml.enc") {
					t.Error("Expected enc.yml.enc to exist")
				}
				content, _ := ioutil.ReadFile("data/enc.yml.enc")
				encFilePattern := "^users:\\n- name: \\w+\\n  pass: ENC\\(.{48}\\)\\n- name: \\w+\\n  pass: ENC\\(.{48}\\)\\n$"
				if !regexp.MustCompile(encFilePattern).Match(content) {
					t.Error("Expected enc.yml.enc to be valid")
				}
			}
		})
	}
}
