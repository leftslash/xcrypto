package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leftslash/xcrypto"
)

func main() {

	key := flag.String("key", "", "key to decrypt message")
	msg := flag.String("msg", "", "message to decrypt")
	flag.Parse()

	if *key == "" {
		fmt.Fprintln(os.Stderr, "key required")
		os.Exit(1)
	}

	if *msg == "" {
		fmt.Fprintln(os.Stderr, "message required")
		os.Exit(1)
	}

	plaintext, err := xcrypto.Decrypt(*key, *msg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", plaintext)
}
