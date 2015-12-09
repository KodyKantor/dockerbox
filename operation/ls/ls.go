// Package ls overrides a normal filesystem 'ls' command.
package ls

import (
	"fmt"
	"strings"

	tm "github.com/buger/goterm"
	"github.com/codegangsta/cli"
	docker "github.com/fsouza/go-dockerclient"
	"stash.veritas.com/scm/kody/dockerbox/utils"
)

// List structure implements opinterface.
type List struct{}

// DoStuff in the ls package will run a 'docker ps' and make the results
// look like a normal filesystem 'ls' command.
func (l *List) DoStuff(c *cli.Context) {

	if len(c.Args()) == 0 {
		// didn't provide a container, so list all of 'em
		utils.RunLinuxCmd(c)
		return
	}

	// user possibly provided a container, so list it's file system
	args := c.Args()
	containerName, err := utils.GetContainerName(args[0])
	if err != nil {
		fmt.Println("Error getting container name:", err)
		return
	}

	if containerName == "" {
		listAllContainers()
		return
	}

	loc := strings.Index(args[0], containerName)
	subDir := args[0][loc+len(containerName):]
	otherArgs := make([]string, 1)
	for _, arg := range args[1:] {
		otherArgs = append(otherArgs, string(arg))
	}
	otherArgs = otherArgs[1:] // cut off the empty string
	//fmt.Printf("other args: %q\n", otherArgs)

	cmd := make([]string, 3)
	cmd[0] = "ls"
	cmd[1] = "/" + subDir
	cmd[2] = "--color=tty"
	cmd = append(cmd, otherArgs...)

	utils.RunCmd(containerName, cmd, false)
}

func listAllContainers() error {

	// Init the client
	client, err := docker.NewClientFromEnv()
	if err != nil {
		fmt.Println("Couldn't connect to docker:", err)
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		return fmt.Errorf("Error listing containers: %v", err)
	}

	table := tm.NewTable(0, 10, 5, ' ', 0)

	for num, c := range containers {
		//fmt.Println("Container:", c.Names[0])
		name := c.Names[0][1:]
		inspection, err := client.InspectContainer(c.ID)
		if err != nil {
			return fmt.Errorf("Error inspecting container: %v", err)
		}

		// if running, color green, else red
		if inspection.State.Running {
			name = tm.Color(name, tm.GREEN)
		} else {
			name = tm.Color(name, tm.RED)
		}

		// print three containers per row
		// this is stupid, but I don't want to spend hours on this junk
		if num%3 == 0 {
			num = 0
			fmt.Fprintln(table)
		}
		fmt.Fprintf(table, "%s\t", name)
	}
	tm.Println(table)
	tm.Flush()

	if err != nil {
		return fmt.Errorf("Error flushing: %v", err)
	}
	return nil
}
