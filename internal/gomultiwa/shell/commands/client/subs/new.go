package subs

import (
	"bytes"

	"github.com/abiosoft/ishell"
	"github.com/mdp/qrterminal"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
)

// GetCmdNew returns the command to create a new client
func GetCmdNew(gmw gmwi.GoMultiWAInterface) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "new",
		Help: "Create a new client",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			uss := []string{}
			uids := []string{}
			for _, u := range *gmw.GetUserManager().Userconfig.Users {
				uss = append(uss, u.Name)
				uids = append(uids, u.ID)
			}
			i := c.MultiChoice(uss, "User?")
			uid := uids[i]
			us := uss[i]
			c.Println("Registering client for " + us)
			qr, _, e := gmw.StartRegistration(uid)
			if e != nil {
				c.Println("Registration failed! Reason: " + e.Error())
				return
			}
			w := new(bytes.Buffer)
			qrterminal.GenerateWithConfig(<-qr, qrterminal.Config{
				Level:     qrterminal.L,
				Writer:    w,
				BlackChar: qrterminal.BLACK,
				WhiteChar: qrterminal.WHITE,
				QuietZone: 0,
			})
			c.Println(w)
			c.Println("User created successfully")
		},
	}
}
