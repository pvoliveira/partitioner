package main

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Println("Secret value: ")
	bytes, err := terminal.ReadPassword(syscall.Stdin)

	if err != nil {
		fmt.Printf("error on try read value: %v", err)
		return
	}

	fmt.Printf("value inputed: %s, total chars: %d\n", string(bytes), len(bytes))
}
