package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args[1] = "-password=securestring"
	os.Args[2] = "-alg=SHA256"
	
	main()
}