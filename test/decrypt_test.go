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

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantPattern string
		wantCode    int
	}{
		{name: "decrypt/opt-key-missing", args: []string{"decrypt"},
			wantPattern: "key is required", wantCode: 1},
		{name: "decrypt/opt-key-invalid", args: []string{"decrypt", "-k", "invalid"},
			wantPattern: "key is required", wantCode: 1},
		{name: "decrypt/opt-ext-invalid", args: []string{"decrypt", "-k", key, "-e", "invalid"},
			wantPattern: "extension is invalid", wantCode: 1},
		{name: "decrypt/args-files-missing", args: []string{"decrypt", "-k", key},
			wantPattern: "missing files", wantCode: 1},
		{name: "decrypt/args-files-expand-nok", args: []string{"decrypt", "-k", key, "data/["},
			wantPattern: "invalid files", wantCode: 1},
		{name: "decrypt/args-files-expand-missing", args: []string{"decrypt", "-k", key, "nowhere/*.ini"},
			wantPattern: "missing files", wantCode: 1},
		{name: "decrypt/args-files-not-regular", args: []string{"decrypt", "-k", key, "data"},
			wantPattern: "not a regular file", wantCode: 1},
		{name: "decrypt/args-files-bad-ext", args: []string{"decrypt", "-k", key, "data/enc.yml"},
			wantPattern: "not a .enc file", wantCode: 1},
		{name: "decrypt/ok", args: []string{"decrypt", "-k", key, "data/dec.yml.enc"},
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
			if fileutil.ExistsRegular("data/dec.yml") {
				_ = os.Remove("data/dec.yml")
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
				if !fileutil.ExistsRegular("data/dec.yml") {
					t.Error("Expected dec.yml to exist")
				}
				content, _ := ioutil.ReadFile("data/dec.yml")
				encFilePattern := "^users:\\n- name: \\w+\\n  pass: bd852f\\n- name: \\w+\\n  pass: 2b13a9\\n$"
				if !regexp.MustCompile(encFilePattern).Match(content) {
					t.Error("Expected dec.yml to be valid")
				}
			}
		})
	}
}
