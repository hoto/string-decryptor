package main

import (
	"fmt"
	"github.com/hoto/string-decryptor/config"
)

func main() {
	config.ParseArgsAndFlags()
	fmt.Printf("algo=%s password=%s salt=%s body=%s",
		config.EncryptionAlgorithm,
		config.EncryptionPassword,
		config.EncryptionSalt,
		config.StringToDecrypt,
	)
}
