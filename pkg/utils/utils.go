package utils

import (
	"regexp"
	"strings"
)

func CheckEmail(s string) bool {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(s) > 254 || !rxEmail.MatchString(s) {
		return false
	}
	return true

}
func CheckEmpty(name, email, password string) bool {
	name = strings.Trim(name, " ")
	email = strings.Trim(email, " ")
	password = strings.Trim(password, " ")

	if len(name) == 0 {
		return false
	}

	if len(email) == 0 {

		return false
	}

	if len(password) == 0 {

		return false
	}

	return true
}
