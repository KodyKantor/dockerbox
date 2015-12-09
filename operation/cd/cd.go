package cd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"stash.veritas.com/scm/kody/dockerbox/utils"
)

// Chdir implements the Operationer interface.
type Chdir struct{}

// DoStuff will create a shell session inside the chosen container.
func (cd *Chdir) DoStuff(c *cli.Context) {
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
		utils.RunLinuxCmd(c) // they ran 'cd /containers', which doesn't make sense to this program
		return
	}

	//TODO we only currently support 'cd <container name>', which will bring a
	// user into the '/' directory in a container. We should support 'cd <name>/path'
	// and open a shell in the path that they specify.
	cmd := make([]string, 1)
	cmd[0] = "sh"
	utils.RunCmd(containerName, cmd, true)

}
