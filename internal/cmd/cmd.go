package cmd

import flag "github.com/spf13/pflag"

// ICommand is an application sub-command.
type ICommand interface {
	// Name is the command name.
	Name() string
	// Description is the command description.
	Description() string
	// Usage is the command usage template.
	Usage() string
	// Examples is a list of command usage examples.
	Examples() []string
	// FlagSet is the set of flags associated to this command.
	FlagSet() *flag.FlagSet
	// HelpRequested returns true if the help has been requested
	// for this command (help flag provided).
	HelpRequested() bool
	// Run executes the command with the given arguments.
	Run(args ...string) error
}

// Command is an ICommand base implementation.
type Command struct {
	name        string
	description string
	usage       string
	examples    []string
	flagSet     *flag.FlagSet
	helpFlag    *bool
}

// NewCommand creates a base command.
func NewCommand(name string, description string, usage string, examples []string) *Command {
	c := &Command{
		name:        name,
		description: description,
		usage:       usage,
		examples:    examples,
		flagSet:     flag.NewFlagSet("name", flag.ContinueOnError),
	}
	c.flagSet.SortFlags = false
	c.helpFlag = c.flagSet.BoolP("help", "h", false, "Show the help message")
	_ = c.flagSet.MarkHidden("help")
	return c
}

// Name returns the command name.
func (c *Command) Name() string {
	return c.name
}

// Description returns the command description.
func (c *Command) Description() string {
	return c.description
}

// Usage returns the command usage template.
func (c *Command) Usage() string {
	return c.usage
}

// Examples returns the list of command usage examples.
func (c *Command) Examples() []string {
	return c.examples
}

// FlagSet returns the set of flags associated to this command.
func (c *Command) FlagSet() *flag.FlagSet {
	return c.flagSet
}

// HelpRequested returns true if the help has been requested
// for this command (help flag provided).
func (c *Command) HelpRequested() bool {
	return c.helpFlag != nil && *c.helpFlag
}
