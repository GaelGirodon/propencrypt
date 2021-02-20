package test

import (
	"gaelgirodon.fr/propencrypt/internal/core/env"
	"testing"
)

const (
	// key is a sample encryption key.
	key = "e2efc125f1cd4df089819271b75b81d5"

	// pattern is a sample pattern.
	pattern = "pass: (.+)"
)

func TestEncrypt(t *testing.T) {
	tests := []commandTestCase{
		{name: "encrypt/opt-key-missing", args: []string{"encrypt"},
			wantPattern: "key is required", wantCode: 1},
		{name: "encrypt/opt-key-invalid", args: []string{"encrypt", "-k", "invalid"},
			wantPattern: "key is required", wantCode: 1},
		{name: "encrypt/opt-key-invalid-despite-env", args: []string{"encrypt", "-k", "invalid"},
			env:         map[string]string{env.Key: key},
			wantPattern: "key is required", wantCode: 1},
		{name: "encrypt/opt-pattern-missing", args: []string{"encrypt", "-k", key},
			wantPattern: "pattern is required", wantCode: 1},
		{name: "encrypt/opt-pattern-no-group", args: []string{"encrypt", "-k", key, "-p", "invalid"},
			wantPattern: "no capturing group", wantCode: 1},
		{name: "encrypt/opt-pattern-invalid", args: []string{"encrypt", "-k", key, "-p", "[p=(.+)"},
			wantPattern: "pattern is invalid", wantCode: 1},
		{name: "encrypt/opt-pattern-invalid-despite-env", args: []string{"encrypt", "-k", key, "-p", "[p=(.+)"},
			env:         map[string]string{env.Key: key, env.Pattern: pattern},
			wantPattern: "pattern is invalid", wantCode: 1},
		{name: "encrypt/opt-ext-invalid", args: []string{"encrypt", "-k", key, "-p", pattern, "-e", "invalid"},
			wantPattern: "extension is invalid", wantCode: 1},
		{name: "encrypt/opt-ext-invalid-despite-env", args: []string{"encrypt", "-k", key, "-p", pattern, "-e", "invalid"},
			env:         map[string]string{env.Key: key, env.Pattern: pattern, env.Ext: ".ok"},
			wantPattern: "extension is invalid", wantCode: 1},
		{name: "encrypt/args-files-missing", args: []string{"encrypt", "-k", key, "-p", pattern},
			wantPattern: "missing files", wantCode: 1},
		{name: "encrypt/args-files-expand-nok", args: []string{"encrypt", "-k", key, "-p", pattern, "data/["},
			wantPattern: "invalid files", wantCode: 1},
		{name: "encrypt/args-files-expand-missing", args: []string{"encrypt", "-k", key, "-p", pattern, "nowhere/*.ini"},
			wantPattern: "missing files", wantCode: 1},
		{name: "encrypt/args-files-not-regular", args: []string{"encrypt", "-k", key, "-p", pattern, "data"},
			wantPattern: "not a regular file", wantCode: 1},
		{name: "encrypt/invalid-match", args: []string{"encrypt", "-k", key, "-p", "(pass): (.+)", "data/ok-enc.yml"},
			wantPattern: "invalid match for the property 0 in file '.+'", wantCode: 1},
		{name: "encrypt/ok-no-match", args: []string{"encrypt", "-k", key, "-p", pattern, "data/warn-no-match.yml"},
			wantPattern: "Warn: no property to process in '.+'", wantCode: 0, wantFile: "data/warn-no-match.yml.enc",
			wantFilePattern: "^users:\n$"},
		{name: "encrypt/ok", args: []string{"encrypt", "-k", key, "-p", pattern, "data/ok-enc.yml"},
			wantPattern: "", wantCode: 0, wantFile: "data/ok-enc.yml.enc",
			wantFilePattern: "^users:\\n- name: \\w+\\n  pass: ENC\\(.{48}\\)\\n- name: \\w+\\n  pass: ENC\\(.{48}\\)\\n$"},
		{name: "encrypt/ok-env", args: []string{"encrypt", "data/ok-enc.yml"},
			env:         map[string]string{env.Key: key, env.Pattern: pattern},
			wantPattern: "", wantCode: 0, wantFile: "data/ok-enc.yml.enc",
			wantFilePattern: "^users:\\n- name: \\w+\\n  pass: ENC\\(.{48}\\)\\n- name: \\w+\\n  pass: ENC\\(.{48}\\)\\n$"},
	}
	// Run command test cases
	testCommand(t, tests, []string{"ok-enc.yml.enc", "warn-no-match.yml.enc"})
}
