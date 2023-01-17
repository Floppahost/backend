package handler

import "fmt"

var UnauthorizedErr = "you don't have permission to perform this action"

func Errors(err error) (int, string) {
	errString := fmt.Sprintf("%v", err)
	switch errString {

	// ---------------- 400 ----------------
	case "email or username already exists":
		return 400, errString
	case "invalid old password":
		return 400, errString
	case "this domain already exists":
		return 400, errString

	// ---------------- 401 ----------------
	case "unauthorized":
		return 401, UnauthorizedErr
	case "you don't have permission to perform this action":
		return 401, errString
	case "the user is currently blacklisted":
		return 401, errString
	case "invalid credentials":
		return 401, errString

		// ---------------- 404 ----------------
	case "the requested user doesn't exist":
		return 404, errString
	case "invalid upload":
		return 404, errString
	case "invalid domain":
		return 404, "this domain doesn't exist"
	case "invalid invite":
		return 404, errString
	case "the user doesn't exist":
		return 404, errString
	}

	return 500, errString
}
