package subs

import (
	"log"

	"github.com/abiosoft/ishell"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
)

// GetCmdNew returns the command to create a new user
func GetCmdNew(gmw gmwi.GoMultiWAInterface) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "new",
		Help: "Create a new user",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			c.Print("Full Name:")
			name := c.ReadLine()
			c.Print("Mail:")
			mail := c.ReadLine()
			c.Print("Password:")
			password := c.ReadPassword()
			c.Print("Password repeated:")
			if pwr := c.ReadPassword(); pwr != password {
				c.Println("Passwords do not match!")
				return
			}
			um := gmw.GetUserManager()

			id, err := um.CreateUser(name, mail)
			if err != nil {
				c.Println(err)
				return
			}
			if err := um.SetUserPW(id, password); err != nil {
				log.Println(err)
			}
			if err := um.SetUserAdmin(id, c.MultiChoice([]string{"no", "yes"}, "admin?") == 1); err != nil {
				log.Println(err)
			}
			c.Println("User created successfully")
		},
	}
}
