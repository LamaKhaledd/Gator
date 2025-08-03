package app

import (
	"fmt"

	"gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.handlers == nil {
		c.handlers = make(map[string]func(*state, command) error)
	}
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}
	username := cmd.args[0]
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}
	fmt.Printf("User set to '%s'\n", username)
	return nil
}

func RunCommand(cfg *config.Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no command provided")
	}

	appState := &state{cfg: cfg}
	cmds := &commands{}
	cmds.register("login", handlerLogin)

	cmd := command{
		name: args[0],
		args: args[1:],
	}

	return cmds.run(appState, cmd)
}
