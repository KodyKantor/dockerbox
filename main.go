package main

import (
	"os"
	"strings"

	"stash.veritas.com/scm/kody/dockerbox/operation"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "gives a devops engineer the power to use the tools they know for the job they love <3"
	app.Name = os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	app.Action = pickFunction

	app.Run(os.Args)
}

func pickFunction(c *cli.Context) {
	operation.GetOperation(c.App.Name).DoStuff(c)
}
