package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"

	"strings"

	"golang.org/x/crypto/argon2"
)

func VerifyPassword(password, encodedHash string) error {
	parts := strings.Split(encodedHash, ".")
	if len(parts) != 2 {
		return ErrorHandler(errors.New("invalid encoded hash format"), "internal server error")
	}

	saltBse64 := parts[0]
	hashedPasswordBase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBse64)
	if err != nil {
		return ErrorHandler(err, "internal server error")
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(hashedPasswordBase64)
	if err != nil {
		return ErrorHandler(err, "internal error")

	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		return ErrorHandler(errors.New("hash length mismatch"), "incorrect password")
	}

	if subtle.ConstantTimeCompare(hash, hashedPassword) == 1 {
		return nil // all other condition other than this will return error
	}
	// return ErrorHandler(errors.New("incorrect password"), "...incorrect password") // This error return indicate that the password provided by the user does not match the stored hashed password. But the comparison process itself was successfull. so there is no other error in the process. Thats why we come to this.
	return ErrorHandler(errors.New("incorrect password"), "Invalid username or password")

}
