package decrypt

import (
	"gaelgirodon.fr/propencrypt/internal/cmd"
	"gaelgirodon.fr/propencrypt/internal/core"
	"gaelgirodon.fr/propencrypt/internal/env"
	"gaelgirodon.fr/propencrypt/internal/validate"
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
			"# Decrypt properties in multiple files (edit them in place)\n" +
				"decrypt -k [...] -e \"\" *.properties",
		})
	return &Command{
		Command: c,
		flags: Flags{
			key: c.FlagSet().StringP("key", "k", "", "256-bit encryption key"),
			ext: c.FlagSet().StringP("ext", "e", ".enc", "Extension to remove from filenames"),
		},
	}
}

// Run decrypts properties in files.
func (c *Command) Run(args ...string) error {
	// Set options values to environment variables values by default
	c.flags.key = c.GetFlagOrEnv("key", c.flags.key, env.Key)
	c.flags.ext = c.GetFlagOrEnv("ext", c.flags.ext, env.Ext)
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
	pattern := regexp.MustCompile(`(ENC\(([\w+/=]+)\))`)
	return core.Decrypt(files, pattern, key, ext)
}
