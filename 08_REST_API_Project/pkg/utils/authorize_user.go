package utils

import "errors"

type ContextKey string

func AuthorizeUser(userRole string, allowedRole ...string) (bool, error) {
	for _, allowedRole := range allowedRole {
		if userRole == allowedRole {
			return true, nil
		}
	}
	return false, errors.New("user not authorized")
}
