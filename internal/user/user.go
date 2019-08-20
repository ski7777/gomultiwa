package user

import (
	wac "github.com/ski7777/gomultiwa/internal/waclient"
)

type User struct {
	clients []wac.WAClient
}

const (
	WHITELIST = 0
	BLACKLIST = 1
)

type UserPermisson struct {
	cleint               wac.WAClient
	whitelist, blacklist []Permission
	mode                 int
}

type Permission struct {
	remote string
}
