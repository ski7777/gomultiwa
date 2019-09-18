package commands

import "github.com/abiosoft/ishell"
import gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"

func GetCmdSave(gmw gmwi.GoMultiWAInterface) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "save",
		Help: "Save config",
		Func: func(c *ishell.Context) {
			gmw.SaveConfig()
		},
	}
}
