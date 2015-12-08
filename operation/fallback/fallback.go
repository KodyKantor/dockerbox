package fallback

import (
	"fmt"
	"strings"

	"stash.veritas.com/scm/kody/dockerbox/utils"

	"github.com/codegangsta/cli"
)

// Fallback implements the Operationer interface, and is used
// when a command wasn't specifically implemented. Results may vary.
type Fallback struct {
	Operation string //the operation that isn't implemented (the one that got us here)
}

// DoStuff takes a cli context, and tries its best to do what the
// user intended.
func (f *Fallback) DoStuff(c *cli.Context) {

	args := c.Args()
	containerName, err := utils.GetContainerName(args[0])
	if err != nil {
		fmt.Println("Error getting container name:", err)
	}

	loc := strings.Index(args[0], containerName)
	subDir := args[0][loc+len(containerName):]
	otherArgs := make([]string, 1)
	for _, arg := range args[1:] {
		otherArgs = append(otherArgs, string(arg))
	}
	otherArgs = otherArgs[1:] // cut off the empty string
	//fmt.Printf("other args: %q\n", otherArgs)

	cmd := make([]string, 2)
	cmd[0] = f.Operation
	cmd[1] = subDir
	cmd = append(cmd, otherArgs...)

	fmt.Printf("Trying: %q\n", cmd)
	utils.RunCmd(containerName, cmd, false)

}
