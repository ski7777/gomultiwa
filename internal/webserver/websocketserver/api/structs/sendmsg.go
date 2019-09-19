package structs

// SendmsgReq represents the requets struct for sendmsg call
type SendmsgReq struct {
	ID      string `json:"id"`
	MSG     string `json:"msg"`
	JID     string `json:"jid"`
	Session string `json:"sess"`
}
