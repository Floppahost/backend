package handler

import "fmt"

func Errors(err error) (int, string) {
	errString := fmt.Sprintf("%v", err)
		switch errString {
		case "you don't have permission to perform this action":
			return 401, errString
		case "the requested user doesn't exist":
			return 404, errString
		case "the user is currently blacklisted":
			return 401, errString
		}
		return 500, errString
}