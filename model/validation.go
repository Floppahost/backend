package model

type UserValidation struct {
	ValidUser   bool
	Admin       bool
	Blacklisted bool
	Username    string
	Uid         int
	JWT         string
	ApiKey      string
}