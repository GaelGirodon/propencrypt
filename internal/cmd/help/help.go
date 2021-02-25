package help

import (
	"errors"
	"gaelgirodon.fr/propencrypt/internal/cmd"
	"gaelgirodon.fr/propencrypt/internal/log"
	flag "github.com/spf13/pflag"
	"strconv"
	"strings"
)

// Command is the help command structure.
type Command struct {
	*cmd.Command
	commands    []cmd.ICommand
	rootFlagSet *flag.FlagSet
}

// NewHelpCmd initializes a help command.
func NewHelpCmd(commands []cmd.ICommand, rootFlagSet *flag.FlagSet) cmd.ICommand {
	return &Command{
		Command:     cmd.NewCommand("help", "", "", []string{}),
		commands:    commands,
		rootFlagSet: rootFlagSet,
	}
}

// Run shows the help message.
func (c *Command) Run(args ...string) (err error) {
	if len(args) == 0 {
		showGlobalHelpMessage(c.commands, c.rootFlagSet)
	} else if len(args) == 1 {
		err = showCommandHelpMessage(c.commands, args[0])
	} else {
		err = errors.New("invalid command")
	}
	return err
}

// showGlobalHelpMessage writes the global help message to the log output.
func showGlobalHelpMessage(commands []cmd.ICommand, rootFlagSet *flag.FlagSet) {
	log.Println("propencrypt encrypts and decrypts properties in files.")
	log.Println("\nUsage:\n  propencrypt <command> [options]")
	log.Println("\nCommands:")
	nameMaxWidth := 0
	for _, command := range commands {
		if len(command.Name()) > nameMaxWidth {
			nameMaxWidth = len(command.Name())
		}
	}
	for _, command := range commands {
		log.Printf("  %-"+strconv.Itoa(nameMaxWidth)+"s  %s\n",
			command.Name(), command.Description())
	}
	log.Println("\nOptions:")
	log.Print(rootFlagSet.FlagUsages())
	log.Println("\nUse \"propencrypt help <command>\" for more information about a given command.")
}

// showCommandHelpMessage writes the help message of a command to the log output.
func showCommandHelpMessage(commands []cmd.ICommand, commandName string) error {
	var selectedCommand cmd.ICommand
	for _, command := range commands {
		if command.Name() == commandName {
			selectedCommand = command
		}
	}
	if selectedCommand == nil {
		return errors.New("unknown command")
	}
	log.Println(selectedCommand.Description())
	log.Println("\nUsage:\n  propencrypt " + selectedCommand.Usage())
	if selectedCommand.FlagSet() != nil {
		log.Println("\nOptions:")
		log.Print(selectedCommand.FlagSet().FlagUsages())
	}
	if len(selectedCommand.Examples()) > 0 {
		log.Print("\nExamples:")
		for _, e := range selectedCommand.Examples() {
			log.Println("\n  " + strings.ReplaceAll(e, "\n", "\n  "))
		}
	}
	return nil
}
