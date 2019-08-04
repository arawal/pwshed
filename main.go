package main

import (
	"flag"
	"fmt"

	"github.com/arawal/pwshed/hashlib"
	"github.com/arawal/pwshed/logger"
	"github.com/arawal/pwshed/server"
	"github.com/arawal/pwshed/stats"
)

func main() {
	stats.Init()
	logger.Init()

	// set up and parse command line flags
	cli := flag.Bool("cli", false, "Launch CLI or the API version")
	password := flag.String("password", "", "Your secure password")
	hashAlg := flag.String("alg", "bcrypt", "hashing algorithm to use")
	flag.Parse()

	if !*cli {
		server.LaunchServer()
	} else {
		err := validateInput(*password, *hashAlg)
		if err != nil {
			logger.Error("cli", err.Error())
			return
		}

		result, err := hashlib.Hash(*password, *hashAlg)
		if err != nil {
			logger.Error("cli", err.Error())
		} else {
			logger.Info("cli", fmt.Sprintf("passwords hashed this session: %d", stats.CurrentStats.Count))
			fmt.Println(result)
		}
		return
	}
}

// validateInput validates input from the cli
/*
	Input:
		- password - string - clear text password to be hashed
		- hashAlg - string - Hashing algorithm to use
	Output:
		- err - error - error encountered when hashing
*/
func validateInput(password, hashAlg string) error {
	// input validation
	if password == "" {
		return fmt.Errorf("missing required -password argument")
	}

	if hashAlg != "MD5" && hashAlg != "SHA256" && hashAlg != "SHA512" && hashAlg != "SHA3" && hashAlg != "bcrypt" && hashAlg != "" {
		return fmt.Errorf("we currently only support SHA256, SHA512, MD5, SHA3 and bcrypt(current standard) algorithms")
	}

	return nil
}
