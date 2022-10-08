package version

import (
	"gaelgirodon.fr/propencrypt/internal/cmd"
	"gaelgirodon.fr/propencrypt/internal/log"
)

// Version is the application version number.
const Version = "0.3.0"

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
func (c *Command) Run(...string) error {
	log.Println("propencrypt version " + Version)
	return nil
}
