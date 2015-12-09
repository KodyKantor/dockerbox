// Package start starts containers.
package start

import (
	"fmt"

	"github.com/codegangsta/cli"
	"stash.veritas.com/scm/kody/dockerbox/utils"
)

// Start implements the Operationer interface.
type Start struct{}

// DoStuff will attempt to start the provided container
func (s *Start) DoStuff(c *cli.Context) {
	if len(c.Args()) == 0 {
		utils.RunLinuxCmd(c)
		return
	}

	args := c.Args()
	containerName, err := utils.GetContainerName(args[0])
	if err != nil {
		fmt.Println("Error getting container name:", err)
		return
	}

	if containerName == "" {
		utils.RunLinuxCmd(c)
		return
	}

}
