package client

import (
	"github.com/abiosoft/ishell"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/shell/commands/client/subs"
)

// GetCmdClient returns the command for editing the client database
func GetCmdClient(gmw gmwi.GoMultiWAInterface) *ishell.Cmd {
	cmd := &ishell.Cmd{
		Name: "client",
		Help: "Edit user database",
		Func: func(c *ishell.Context) {
			c.Println(c.Cmd.HelpText())
		},
	}
	cmd.AddCmd(subs.GetCmdNew(gmw))
	return cmd
}
