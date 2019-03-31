package command

import "fmt"

// Command is type for tower's command
type Command interface {
	Name() string
	Describe() string
	Usage() string
	Handle() error
}

// commands store all tower's command
var commands map[string]Command

// Register register new command to tower
func Register(cmd Command) error {
	if _, ok := commands[cmd.Name()]; ok {
		return fmt.Errorf("cmd %s used", cmd.Name())
	}
	commands[cmd.Name()] = cmd
	return nil
}

// Get return the command searched by arg
func Get(cmdStr string) Command {
	if cmd, ok := commands[cmdStr]; ok {
		return cmd
	}
	return nil
}

func init() {
	commands = make(map[string]Command)
}
