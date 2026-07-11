package tools

import "time"

type Command struct {
	Name    string
	Args    []string
	Dir     string
	Env     map[string]string
	Timeout time.Duration
}

type CommandOption func(*Command)

func WithArgs(args ...string) CommandOption {
	return func(c *Command) { c.Args = args }
}

func WithDir(dir string) CommandOption {
	return func(c *Command) { c.Dir = dir }
}

func WithEnv(env map[string]string) CommandOption {
	return func(c *Command) { c.Env = env }
}

func WithTimeout(d time.Duration) CommandOption {
	return func(c *Command) {
		if d > 0 {
			c.Timeout = d
		}
	}
}

func NewCommand(name string, opts ...CommandOption) *Command {
	cmd := &Command{Name: name, Timeout: 5 * time.Minute}
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}
