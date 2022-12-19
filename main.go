package main

import (
	"fmt"
	"github.com/hoto/string-decryptor/config"
	"github.com/logrusorgru/aurora"
)

func main() {
	config.ParseArgsAndFlags()
	fmt.Printf("hello %s\n", aurora.Cyan("world"))
}
