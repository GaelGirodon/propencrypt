package propencrypt

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gaelgirodon.fr/propencrypt/internal/core/log"
	"gaelgirodon.fr/propencrypt/pkg/crypto"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Encrypt encrypts property values in a list of file.
func Encrypt(filenames []string, pattern *regexp.Regexp, key *[32]byte, ext string) error {
	return process(true, filenames, pattern, key, ext)
}

// Decrypt decrypts property values in a list of file.
func Decrypt(filenames []string, pattern *regexp.Regexp, key *[32]byte, ext string) error {
	return process(false, filenames, pattern, key, ext)
}

// process encrypts (encrypt) or decrypts (!encrypt) property values in a list of file.
func process(encrypt bool, filenames []string, pattern *regexp.Regexp, key *[32]byte, ext string) error {
	for _, f := range filenames {
		// Get file info
		fileInfo, err := os.Stat(f)
		if err != nil {
			return errors.New(fmt.Sprintf("unable to get information about the file '%s", f))
		}
		// Read the file content
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return errors.New(fmt.Sprintf("unable to read file '%s", f))
		}
		// Find all matches of properties to encrypt
		matches := pattern.FindAllSubmatchIndex(content, -1)
		if len(matches) == 0 {
			log.Warn("no property to process in  '%s'", f)
		}
		// Encrypt each property and generate the output file content
		var output []byte
		offset := 0
		for i, match := range matches {
			if encrypt && len(match) != 4 || !encrypt && len(match) != 6 {
				return errors.New(fmt.Sprintf("invalid match for the #%d property in file '%s", i, f))
			}
			// Append unprocessed content to the output
			output = append(output, content[offset:match[2]]...)
			// Encrypt or decrypt the value
			var processedValue []byte
			if encrypt {
				encryptedValue, err := crypto.Encrypt(content[match[2]:match[3]], key)
				if err != nil {
					return errors.New(fmt.Sprintf("unable to encrypt the #%d property in file '%s", i, f))
				}
				// Encode the encrypted value and append it to the output
				processedValue = []byte(fmt.Sprintf("ENC(%s)", base64.StdEncoding.EncodeToString(encryptedValue)))
			} else {
				valueToDecrypt, err := base64.StdEncoding.DecodeString(string(content[match[4]:match[5]]))
				if err != nil {
					return errors.New(fmt.Sprintf("unable to decode the #%d property in file '%s", i, f))
				}
				processedValue, err = crypto.Decrypt(valueToDecrypt, key)
				if err != nil {
					return errors.New(fmt.Sprintf("unable to decrypt the #%d property in file '%s", i, f))
				}
			}
			output = append(output, processedValue...)
			offset = match[3]
		}
		// Append remaining unprocessed content
		output = append(output, content[offset:]...)
		// Write the output file
		var outputFilename string
		if encrypt {
			outputFilename = f + ext
		} else {
			outputFilename = strings.TrimSuffix(f, ext)
		}
		if err = ioutil.WriteFile(outputFilename, output, fileInfo.Mode().Perm()); err != nil {
			return errors.New(fmt.Sprintf("unable to write output file '%s", outputFilename))
		}
	}
	return nil
}
