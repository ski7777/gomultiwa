package config

import (
	"github.com/abiosoft/ishell"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/shell/commands/config/subs"
)

// GetCmdConfig returns the command for editing the user config
func GetCmdConfig(gmw gmwi.GoMultiWAInterface) *ishell.Cmd {
	cmd := &ishell.Cmd{
		Name: "config",
		Help: "Edit config",
		Func: func(c *ishell.Context) {
			c.Println(c.Cmd.HelpText())
		},
	}
	cmd.AddCmd(subs.GetCmdSave(gmw))
	return cmd
}
