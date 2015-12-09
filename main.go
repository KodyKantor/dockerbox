package main

import (
	"os"
	"strings"

	"stash.veritas.com/scm/kody/dockerbox/operation"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "present docker containers as a semi-posix filesystem!"
	app.Name = os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	app.Action = pickFunction

	app.Run(os.Args)
}

func pickFunction(c *cli.Context) {
	operation.GetOperation(c.App.Name).DoStuff(c)
}
