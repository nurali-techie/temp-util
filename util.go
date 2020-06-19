package util

import (
	"fmt"
	"strings"
)

// UserName joins given firstName and lastName and return string in UserName format
func UserName(firstName, lastName string) string {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)

	return fmt.Sprintf("%s-%s", firstName, lastName)
}
