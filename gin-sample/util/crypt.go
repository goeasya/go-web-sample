package util

import (
	"crypto/sha256"
	"fmt"
)

const keySalt = "7QA9ab3cd%f012*6"

func EnCrypt(password string) string {
	commonKey := keySalt + password
	retPassword := fmt.Sprintf("%x", sha256.Sum224([]byte(commonKey)))
	return retPassword[0:16]
}

func ValidatePassword(actual, except string) bool {
	return EnCrypt(actual) == except
}
