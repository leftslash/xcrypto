package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leftslash/xcrypto"
)

func main() {

	key := flag.String("key", "", "key to encrypt message")
	msg := flag.String("msg", "", "message to encrypt")
	flag.Parse()

	if *key == "" {
		fmt.Fprintln(os.Stderr, "key required")
		os.Exit(1)
	}

	if *msg == "" {
		fmt.Fprintln(os.Stderr, "message required")
		os.Exit(1)
	}

	ciphertext, err := xcrypto.Encrypt(*key, *msg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", ciphertext)
}
