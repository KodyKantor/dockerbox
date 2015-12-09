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
	app.Authors = []cli.Author{cli.Author{Name: "Kody A Kantor", Email: "kody.kantor@veritas.com"}}
	app.Version = "Hackathon 2.0"

	app.Run(os.Args)
}

// pickFunction will find out which subcommand (if any) is applicable
// to the name of the binary that was called. It then calls the DoStuff function
// on that operation (as per the operationer interface).
func pickFunction(c *cli.Context) {
	operation.GetOperation(c.App.Name).DoStuff(c)
}
