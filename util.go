package util

import (
	"fmt"
	"strings"
)

// UserName joins given firstName and lastName and return string in UserName format
func UserName(firstName, lastName string) string {
	name := fmt.Sprintf("%s-%s", strings.ToLower(firstName), strings.ToLower(lastName))
	return name
}
