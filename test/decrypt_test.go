package test

import (
	"gaelgirodon.fr/propencrypt/internal/core/env"
	"testing"
)

func TestDecrypt(t *testing.T) {
	tests := []commandTestCase{
		{name: "decrypt/opt-key-missing", args: []string{"decrypt"},
			wantPattern: "key is required", wantCode: 1},
		{name: "decrypt/opt-key-invalid", args: []string{"decrypt", "-k", "invalid"},
			wantPattern: "key is required", wantCode: 1},
		{name: "decrypt/opt-key-invalid-despite-env", args: []string{"decrypt", "-k", "invalid"},
			env:         map[string]string{env.Key: key},
			wantPattern: "key is required", wantCode: 1},
		{name: "decrypt/opt-ext-invalid", args: []string{"decrypt", "-k", key, "-e", "invalid"},
			wantPattern: "extension is invalid", wantCode: 1},
		{name: "decrypt/opt-ext-invalid-despite-env", args: []string{"decrypt", "-k", key, "-e", "invalid"},
			env:         map[string]string{env.Key: key, env.Ext: ".ok"},
			wantPattern: "extension is invalid", wantCode: 1},
		{name: "decrypt/args-files-missing", args: []string{"decrypt", "-k", key},
			wantPattern: "missing files", wantCode: 1},
		{name: "decrypt/args-files-expand-nok", args: []string{"decrypt", "-k", key, "data/["},
			wantPattern: "invalid files", wantCode: 1},
		{name: "decrypt/args-files-expand-missing", args: []string{"decrypt", "-k", key, "nowhere/*.ini"},
			wantPattern: "missing files", wantCode: 1},
		{name: "decrypt/args-files-not-regular", args: []string{"decrypt", "-k", key, "data"},
			wantPattern: "not a regular file", wantCode: 1},
		{name: "decrypt/args-files-bad-ext", args: []string{"decrypt", "-k", key, "data/ok-enc.yml"},
			wantPattern: "not a .enc file", wantCode: 1},
		{name: "decrypt/err-decode", args: []string{"decrypt", "-k", key, "data/err-decode.yml.enc"},
			wantPattern: "unable to decode the property 0 in file '.+'", wantCode: 1},
		{name: "decrypt/err-decrypt", args: []string{"decrypt", "-k", key, "data/err-decrypt.yml.enc"},
			wantPattern: "unable to decrypt the property 0 in file '.+'", wantCode: 1},
		{name: "decrypt/ok", args: []string{"decrypt", "-k", key, "data/ok-dec.yml.enc"},
			wantPattern: "", wantCode: 0, wantFile: "data/ok-dec.yml",
			wantFilePattern: "^users:\\n- name: \\w+\\n  pass: bd852f\\n- name: \\w+\\n  pass: 2b13a9\\n$"},
		{name: "decrypt/ok-env", args: []string{"decrypt", "data/ok-dec.yml.enc"},
			env:         map[string]string{env.Key: key},
			wantPattern: "", wantCode: 0, wantFile: "data/ok-dec.yml",
			wantFilePattern: "^users:\\n- name: \\w+\\n  pass: bd852f\\n- name: \\w+\\n  pass: 2b13a9\\n$"},
	}
	// Run command test cases
	testCommand(t, tests, []string{"ok-dec.yml"})
}
