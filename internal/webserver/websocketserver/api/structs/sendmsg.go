package structs

type SendmsgReq struct {
	ID  string `json:"id"`
	MSG string `json:"msg"`
	JID string `json:"jid"`
}
