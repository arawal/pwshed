package main

import (
	"flag"
	"fmt"

	"github.com/arawal/pwshed/hashlib"
)

func main() {
	// Init config
	// Get user input
	// expects: password, hash algorithm (optional)
	password := flag.String("password", "", "Your secure password")
	hashAlg := flag.String("alg", "SHA3", "hashing algorithm to use")

	flag.Parse()

	if *password == "" {
		fmt.Println("missing required -password argument")
		return
	}

	if *hashAlg != "MD5" && *hashAlg != "SHA256" && *hashAlg != "SHA512" && *hashAlg != "SHA3" && *hashAlg != "" {
		fmt.Println("we currently only support SHA256, SHA512, MD5 and SHA3 algorithms")
		return
	}

	// Hash input
	result, err := hashlib.Hash(*password, *hashAlg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	return
}
