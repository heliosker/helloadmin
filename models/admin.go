package models

func CheckAdminUser(username, password string) bool {
	if username == "admin" && password == "123456" {
		return true
	}
	return false
}
