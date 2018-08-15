package main

import (
	"github.com/nicholaswagner/visca_cli"
	"fmt"
)

var Build string // when compile pass:  -ldflags "-X main.Build <build sha1>"


func main() {
	fmt.Printf("Using build: %s\n", Build)
}