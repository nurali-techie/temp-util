package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// UserName joins given firstName, lastName, middleName and return string in UserName format
func UserName(firstName, lastName, middleName string) string {
	firstName = strings.ToLower(strings.TrimSpace(firstName))
	lastName = strings.ToLower(strings.TrimSpace(lastName))
	middleName = strings.ToLower(strings.TrimSpace(middleName))

	initial := string(middleName[0])
	return fmt.Sprintf("%s-%s-%s", firstName, lastName, initial)
}

func GeneratePassword() string {
	return strconv.Itoa(time.Now().Nanosecond())
}

func GenerateEmail(userName string) string {
	userName = strings.ReplaceAll(userName, "-", ".")
	return fmt.Sprintf("%s@mycompany.com", userName)
}
