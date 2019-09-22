package shell

import (
	"github.com/abiosoft/ishell"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/shell/commands"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/shell/commands/config"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/shell/commands/user"
)

// Shell represents the gomultiwa shell
type Shell struct {
	gmw   gmwi.GoMultiWAInterface
	Shell *ishell.Shell
}

// NewShell returns a new Schell struct
func NewShell(gmw gmwi.GoMultiWAInterface) *Shell {
	s := new(Shell)
	s.Shell = ishell.New()
	s.gmw = gmw
	s.Shell.Interrupt(interruptFunc(s.gmw))
	s.Shell.AddCmd(config.GetCmdConfig(s.gmw))
	s.Shell.AddCmd(commands.GetCmdExit(s.gmw))
	s.Shell.AddCmd(user.GetCmdUser(s.gmw))
	return s
}

// Start starts the shell
func (s *Shell) Start() {
	s.Shell.Run()
}

func interruptFunc(gmw gmwi.GoMultiWAInterface) func(*ishell.Context, int, string) {
	return func(c *ishell.Context, count int, line string) {
		if count >= 2 {
			gmw.Stop()
		}
		c.Println("Input Ctrl-c once more to exit")
	}
}
