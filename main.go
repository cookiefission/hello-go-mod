package main

import (
	"fmt"
	"os"

	"github.com/CrowdSurge/banner"
	lol "github.com/kris-nova/lolgopher"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "hello-go-mod",
	Short: "Something",
	Long:  `Something`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s %s", date, version, commit)

		w := lol.NewLolWriter()
		w.Write([]byte(banner.PrintS("hello go modules")))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
