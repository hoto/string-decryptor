package config

import (
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
)

var (
	Version     string
	ShortCommit string
	BuildDate   string
	Debug       bool

	EncryptionAlgorithm string
	EncryptionPassword  string
	EncryptionSalt      string
	StringToDecrypt     string
)

const (
	empty                   = ""
	encryptionAlgorithmDesc = "(required) Encryption algorithm"
	encryptionPasswordDesc  = "(required) Encryption password"
	encryptionSaltDesc      = "(required) Encryption salt"
	stringToDecryptDesc     = "(required) String to be decrypted"
	usage                   = `Usage example:

  string-decryptor \
    -a aes256 \
    -p password \
    -s salt \
    -b bc5b3bd34ac8

Flags:

`
)

func ParseArgsAndFlags() {
	flag.Usage = overrideUsage()

	flag.BoolVar(&Debug, "debug", false, "Show verbose debug information")
	showVersion := flag.Bool("version", false, "Show version")
	flag.StringVar(&EncryptionAlgorithm, "a", empty, encryptionAlgorithmDesc)
	flag.StringVar(&EncryptionPassword, "p", empty, encryptionPasswordDesc)
	flag.StringVar(&EncryptionSalt, "s", empty, encryptionSaltDesc)
	flag.StringVar(&StringToDecrypt, "b", empty, stringToDecryptDesc)

	flag.Parse()

	if *showVersion {
		fmt.Printf("string-decryptor version %s, commit %s, build %s\n",
			Version, ShortCommit, BuildDate)
		os.Exit(0)
	}

	if Debug {
		debugLog()
	}
}

func overrideUsage() func() {
	return func() {
		_, _ = fmt.Fprintf(os.Stdout, usage)
		flag.PrintDefaults()
	}
}

func debugLog() {
	fmt.Println("Args:")
	fmt.Printf("  args=%s\n", aurora.Cyan(flag.Args()))
}
