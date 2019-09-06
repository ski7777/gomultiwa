package user

type User struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Mail        string          `json:"mail"`
	Password    string          `json:"password"`
	Clients     []string        `json:"clients"`
	Permissions []UserPermisson `json:"permissions"`
	Admin       bool            `json:"admin"`
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
