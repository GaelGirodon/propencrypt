package version

import (
	"gaelgirodon.fr/propencrypt/internal/cmd"
	"gaelgirodon.fr/propencrypt/internal/core/log"
)

// Version is the application version number.
const Version = "0.1.0"

// Command is the version command structure.
type Command struct {
	*cmd.Command
}

// NewVersionCmd initializes a version command.
func NewVersionCmd() cmd.ICommand {
	return &Command{
		Command: cmd.NewCommand("version", "", "", []string{}),
	}
}

// Run shows the application version number.
func (c *Command) Run(args ...string) error {
	log.Println("propencrypt version " + Version)
	return nil
}
