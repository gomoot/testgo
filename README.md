# testgo
testgo


go build -o bin/main main.go
go build -ldflags "-X main.Version=$VERSION"

package main

import (
  "fmt"
  "os"
)

var version string

func main() {
  fmt.Println(version)
}