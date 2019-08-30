package user

type User struct {
	Clients     []string        `json:"clients"`
	Permissions []UserPermisson `json:"permissions"`
}

const (
	WHITELIST = 0
	BLACKLIST = 1
)

type UserPermisson struct {
	Cleint    string       `json:"client"`
	Whitelist []Permission `json:"whitelist"`
	Blacklist []Permission `json:"blacklist"`
	Mode      int          `json:"mode"`
}

type Permission struct {
	Remote string `json:"remote"`
	Jids   string `json:"jids"`
}
