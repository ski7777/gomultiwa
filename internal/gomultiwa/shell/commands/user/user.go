package user

import (
	"github.com/abiosoft/ishell"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/shell/commands/user/subs"
)

// GetCmdUser returns the command for editing the user database
func GetCmdUser(gmw gmwi.GoMultiWAInterface) *ishell.Cmd {
	cmd := &ishell.Cmd{
		Name: "user",
		Help: "Edit user database",
		Func: func(c *ishell.Context) {
			c.Println(c.Cmd.HelpText())
		},
	}
	cmd.AddCmd(subs.GetCmdNew(gmw))
	return cmd
}
