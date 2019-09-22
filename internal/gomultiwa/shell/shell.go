package shell

import (
	"github.com/abiosoft/ishell"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/shell/commands"
)

// Shell represents the gomultiwa shell
type Shell struct {
	gmw   gmwi.GoMultiWAInterface
	shell *ishell.Shell
}

// NewShell returns a new Schell struct
func NewShell(gmw gmwi.GoMultiWAInterface) *Shell {
	s := new(Shell)
	s.shell = ishell.New()
	s.gmw = gmw
	s.shell.AddCmd(commands.GetCmdSave(gmw))
	s.shell.AddCmd(commands.GetCmdExit(gmw))
	s.shell.AddCmd(commands.GetCmdNewUser(gmw))
	return s
}

// Start starts the shell
func (s *Shell) Start() {
	s.shell.Run()
}
