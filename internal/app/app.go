package app

import (
	"gaelgirodon.fr/propencrypt/internal/cmd"
	"gaelgirodon.fr/propencrypt/internal/cmd/decrypt"
	"gaelgirodon.fr/propencrypt/internal/cmd/encrypt"
	"gaelgirodon.fr/propencrypt/internal/cmd/help"
	"gaelgirodon.fr/propencrypt/internal/cmd/version"
	"gaelgirodon.fr/propencrypt/internal/core/log"
	flag "github.com/spf13/pflag"
	"os"
	"strings"
)

// Run executes the application.
func Run() int {
	// Define commands
	commands := make(map[string]cmd.ICommand)
	commands["encrypt"] = encrypt.NewEncryptCmd()
	commands["decrypt"] = decrypt.NewDecryptCmd()
	commands["help"] = help.NewHelpCmd([]cmd.ICommand{commands["encrypt"], commands["decrypt"]})
	commands["version"] = version.NewVersionCmd()

	// Define root flags
	flagSet := flag.NewFlagSet("root", flag.ContinueOnError)
	helpFlag := flagSet.BoolP("help", "h", false, "Show the help message")
	versionFlag := flagSet.BoolP("version", "v", false, "Show the version number")

	// Show the global help message if there is no argument
	if len(os.Args) < 2 {
		_ = commands["help"].Run()
		return 0
	}

	// Select the command from the first argument
	if _, exists := commands[os.Args[1]]; !exists {
		if len(os.Args) > 2 || !strings.HasPrefix(os.Args[1], "-") {
			return log.Error("unknown command")
		}
		// Not a command: parse global flags (ignore errors: global flag set has ExitOnError)
		_ = flagSet.Parse(os.Args[1:])
		if *helpFlag {
			_ = commands["help"].Run()
			return 0
		} else if *versionFlag {
			_ = commands["version"].Run()
			return 0
		} else {
			return log.Error("unknown command")
		}
	}
	selectedCmd := commands[os.Args[1]]

	// Parse the selected command
	args := os.Args[2:]
	if selectedCmd.FlagSet() != nil {
		err := selectedCmd.FlagSet().Parse(args)
		args = selectedCmd.FlagSet().Args()
		if err != nil {
			return log.Error("invalid command")
		}
	}

	// Show the help message for this command if requested
	if selectedCmd.HelpRequested() && selectedCmd != commands["help"] {
		_ = commands["help"].Run(selectedCmd.Name())
		return 0
	}

	// Run it
	err := selectedCmd.Run(args...)
	if err != nil {
		return log.Error(err.Error())
	}
	return 0
}
