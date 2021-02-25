package encrypt

import (
	"gaelgirodon.fr/propencrypt/internal/cmd"
	"gaelgirodon.fr/propencrypt/internal/core"
	"gaelgirodon.fr/propencrypt/internal/env"
	"gaelgirodon.fr/propencrypt/internal/validate"
)

// Command is the encrypt command structure.
type Command struct {
	*cmd.Command
	flags Flags
}

// Flags are encrypt command flags.
type Flags struct {
	key     *string
	pattern *string
	ext     *string
}

// NewEncryptCmd initializes an encrypt command.
func NewEncryptCmd() cmd.ICommand {
	c := cmd.NewCommand(
		"encrypt", "Encrypt properties in files.",
		"encrypt -k <key> -p <pattern> [-e <ext>] <files>",
		[]string{
			"# Encrypt properties in a single file (-> config.yml.enc)\n" +
				"encrypt -k [...] -p \"password: (.+)\" config.yml",
			"# Encrypt properties in multiple files (edit them in place)\n" +
				"encrypt -k [...] -p \"secret=(.+)\" -e \"\" *.properties",
		})
	return &Command{
		Command: c,
		flags: Flags{
			key:     c.FlagSet().StringP("key", "k", "", "256-bit encryption key"),
			pattern: c.FlagSet().StringP("pattern", "p", "", "Sensitive property pattern"),
			ext:     c.FlagSet().StringP("ext", "e", ".enc", "Extension to append to filenames"),
		},
	}
}

// Run encrypts properties in files.
func (c *Command) Run(args ...string) (err error) {
	// Set options values to environment variables values by default
	c.flags.key = c.GetFlagOrEnv("key", c.flags.key, env.Key)
	c.flags.pattern = c.GetFlagOrEnv("pattern", c.flags.pattern, env.Pattern)
	c.flags.ext = c.GetFlagOrEnv("ext", c.flags.ext, env.Ext)
	// Validate options
	key, err := validate.Key(c.flags.key)
	if err != nil {
		return err
	}
	pattern, err := validate.Pattern(c.flags.pattern)
	if err != nil {
		return err
	}
	ext, err := validate.Ext(c.flags.ext)
	if err != nil {
		return err
	}
	// Validate arguments
	files, err := validate.Files(args, "")
	if err != nil {
		return err
	}
	// Encrypt
	return core.Encrypt(files, pattern, key, ext)
}
