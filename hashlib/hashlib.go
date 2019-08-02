package hashlib

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/sha3"
)

// Hash hashes a password using SHA-3 or a user-preferred algorithm
/*
	Input:
		- password - string - clear text password to be hashed
		- algorithm - string - Hashing algorithm to use
	Output:
		- pwshed - string - base64 encoded hash of the password
		- err - error - error encountered when hashing
*/
func Hash(password, algorithm string) (string, error) {
	if password == "" {
		return "", errors.New("no password provided in request")
	}
	switch algorithm {
	case "SHA256":
		return hashSHA256(password)
	case "SHA512":
		return hashSHA512(password)
	case "MD5":
		return hashMD5(password)
	default:
		return hashSHA3(password)
	}
}

func hashSHA256(password string) (string, error) {
	hasher := sha512.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hashed := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashed), nil
}

func hashSHA512(password string) (string, error) {
	hasher := sha512.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hashed := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashed), nil
}

func hashMD5(password string) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hashed := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashed), nil
}

func hashSHA3(password string) (string, error) {
	hasher := sha3.New512()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hashed := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashed), nil
}
