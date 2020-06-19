package util

import (
	"fmt"
	"strings"
)

func UserName(firstName, lastName string) string {
	name := fmt.Sprintf("%s-%s", strings.ToLower(firstName), strings.ToLower(lastName))
	return name
}
