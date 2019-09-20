package commands

import "github.com/abiosoft/ishell"
import gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"

// GetCmdSave returns the command to save the config
func GetCmdSave(gmw gmwi.GoMultiWAInterface) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "save",
		Help: "Save config",
		Func: func(c *ishell.Context) {
			if err:=gmw.SaveConfig();err!=nil{
				c.Println(err)
			}
		},
	}
}
