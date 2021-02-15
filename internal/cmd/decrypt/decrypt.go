package decrypt

import (
	"gaelgirodon.fr/propencrypt/internal/cmd"
	"gaelgirodon.fr/propencrypt/internal/core/propencrypt"
	"gaelgirodon.fr/propencrypt/internal/core/validate"
	"regexp"
)

// Command is the decrypt command structure.
type Command struct {
	*cmd.Command
	flags Flags
}

// Flags are decrypt command flags.
type Flags struct {
	key *string
	ext *string
}

// NewDecryptCmd initializes a decrypt command.
func NewDecryptCmd() cmd.ICommand {
	c := cmd.NewCommand(
		"decrypt", "Decrypt properties in files.",
		"decrypt -k <key> [-e <ext>] <files>",
		[]string{
			"# Decrypt properties in a single file (-> config.yml)\n" +
				"decrypt -k [...] config.yml.enc",
			"# Decrypt properties in multiple files (edit files in place)\n" +
				"decrypt -k [...] -e \"\" *.properties",
		})
	return &Command{
		Command: c,
		flags: Flags{
			key: c.FlagSet().StringP("key", "k", "", "256-bit encryption key"),
			ext: c.FlagSet().StringP("ext", "e", ".enc", "extension to remove from filenames"),
		},
	}
}

// Run decrypts properties in files.
func (c *Command) Run(args ...string) error {
	// Validate options
	key, err := validate.Key(c.flags.key)
	if err != nil {
		return err
	}
	ext, err := validate.Ext(c.flags.ext)
	if err != nil {
		return err
	}
	// Validate arguments
	files, err := validate.Files(args, ext)
	if err != nil {
		return err
	}
	// Decrypt
	pattern := regexp.MustCompile("(ENC\\(([\\w+/=]+)\\))")
	return propencrypt.Decrypt(files, pattern, key, ext)
}