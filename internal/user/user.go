package user

type User struct {
	clients []string
}

const (
	WHITELIST = 0
	BLACKLIST = 1
)

type UserPermisson struct {
	cleint               string
	whitelist, blacklist []Permission
	mode                 int
}

type Permission struct {
	remote string
}
