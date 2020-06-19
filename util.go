package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// UserName joins given firstName, lastName, middleName and return string in UserName format
func UserName(firstName, lastName, middleName string) string {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	middleName = strings.TrimSpace(middleName)

	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)
	middleName = strings.ToLower(middleName)

	initial := middleName[0]
	return fmt.Sprintf("%s-%s-%c", firstName, lastName, initial)
}

func GeneratePassword() string {
	return strconv.Itoa(rand.Intn(1000))
}

func GenerateEmail(userName string) string {
	userName = strings.ReplaceAll(userName, "-", ".")
	return fmt.Sprintf("%s@mycompany.com", userName)
}
