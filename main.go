package main

import (
	"fmt"

	"github.com/CrowdSurge/banner"
	lol "github.com/kris-nova/lolgopher"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("%s %s %s", date, version, commit)

	w := lol.NewLolWriter()
	w.Write([]byte(banner.PrintS("hello go modules")))
}
