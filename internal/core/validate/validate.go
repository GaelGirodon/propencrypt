package validate

import (
	"errors"
	"gaelgirodon.fr/propencrypt/pkg/fileutil"
	"regexp"
	"strings"
)

var (
	// patternFlagRegex is the regex used to validate the pattern value.
	patternFlagRegex = regexp.MustCompile("\\(.+\\)")

	// extFlagRegex is the regex used to validate the extension value.
	extFlagRegex = regexp.MustCompile("^\\.[\\w-.]+$")
)

// Key validates a key value and returns the key in the right format.
func Key(input *string) (*[32]byte, error) {
	if input == nil || len(strings.TrimSpace(*input)) != 32 {
		return nil, errors.New("a valid 256-bit encryption key is required")
	}
	key := [32]byte{}
	copy(key[:], *input)
	return &key, nil
}

// Pattern validates a pattern value and returns the regex.
func Pattern(input *string) (pattern *regexp.Regexp, err error) {
	if input == nil || len(*input) == 0 {
		return nil, errors.New("the sensitive property pattern is required")
	}
	if matches := patternFlagRegex.FindStringIndex(*input); matches == nil {
		return nil, errors.New("the sensitive property pattern seems to contain no capturing group")
	}
	if pattern, err = regexp.Compile(*input); err != nil {
		return nil, errors.New("the sensitive property pattern is invalid")
	}
	return pattern, nil
}

// Ext validates an extension value and returns the extension.
func Ext(input *string) (string, error) {
	if input != nil && len(*input) > 0 && !extFlagRegex.MatchString(*input) {
		return "", errors.New("the extension is invalid")
	}
	if input == nil {
		return "", nil
	}
	return *input, nil
}

// Files validates a files list and expand globs. ext is the suffix
// each file name is expected to have (leave empty to skip this validation).
func Files(input []string, ext string) ([]string, error) {
	files, err := fileutil.Expand(input)
	if err != nil {
		return nil, errors.New("invalid files: " + err.Error())
	}
	if len(files) == 0 {
		return nil, errors.New("missing files arguments")
	}
	for _, file := range files {
		if !fileutil.ExistsRegular(file) {
			return nil, errors.New("'" + file + "' is not a regular file")
		}
		if len(ext) > 0 && !strings.HasSuffix(file, ext) {
			return nil, errors.New("'" + file + "' is not a " + ext + " file")
		}
	}
	return files, nil
}
