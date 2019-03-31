package help

import (
	"flag"

	"github.com/tower/command"
)

// CMD is struct for help command
type CMD struct{}

// Name return the name of help command
func (*CMD) Name() string {
	return "help"
}

// Describe return the description of help command
func (*CMD) Describe() string {
	return "show usage of command"
}

// Handle handle the help command
func (*CMD) Handle() error {
	args := flag.Args()
	if len(args) < 2 {
		println("command not found")
		return nil
	}
	cmd := command.Get(args[1])
	if cmd != nil {
		cmd.Usage()
	}
	return nil
}

// Usage return the usage of help command
func (*CMD) Usage() string {
	return "help [command]"
}

func init() {
	println("register help command")
	command.Register(&CMD{})
}
