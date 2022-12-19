[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build status](https://github.com/hoto/string-decryptor/workflows/Build%20and%20test/badge.svg?branch=master)](https://github.com/hoto/string-decryptor/actions)
[![Release](https://img.shields.io/github/release/hoto/string-decryptor.svg?style=flat-square)](https://github.com/hoto/string-decryptor/releases/latest)
[![Powered By: goreleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser/goreleaser)

# String decryptor

CLI tool to decrypt passed in string.

### Installation
    
Mac:

    brew install hoto/homebrew-repo/string-decryptor

Mac or Linux:

    sudo curl -L \
      "https://github.com/hoto/string-decryptor/releases/download/1.0.0/string-decryptor_1.0.0_$(uname -s)_$(uname -m)" \
       -o /usr/local/bin/string-decryptor

    sudo chmod +x /usr/local/bin/string-decryptor
    
Or manually download binary from [releases](https://github.com/hoto/string-decryptor/releases).
    
### Run

    string-decryptor --help
    string-decryptor --version
    string-decryptor
    
### Development

Build and test:

    go get github.com/hoto/string-decryptor
    
    make dependencies build test
    
Build binary:

     make build
    ./bin/string-decryptor

Run with arguments:

    make args="myargs" run

Install to global golang bin directory:

    make install
    string-decryptor
    
### Release

Add a git tag and push it to GitHub.  
GitHub action will use `goreleaser` to automatically releases formula to homebrew repo at [hoto/homebrew-repo](https://github.com/hoto/homebrew-repo) on every git tag.
