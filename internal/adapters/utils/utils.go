package utils

import "strings"

func CheckEmail(s string) bool {
	for _, i := range s {
		if i == '@' {
			for _, j := range s {
				if j == '.' {

					return true
				}
			}

		}
	}

	return false
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
