// Package utils provides functions that all operations can use
// to kick off the execution of their commands.
package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
)

// RunCmd takes the name of a container, and a command to run (a list of arguments).
// It sets up a 'docker exec'-type command and executes it.
func RunCmd(container string, command []string, stdin bool) {
	//fmt.Printf("Trying %q\n", command)
	client, err := docker.NewClientFromEnv()
	if err != nil {
		fmt.Println("Couldn't connect to docker:", err)
		return
	}

	//fmt.Println("container:", container)
	// allocate an execution environment
	exec, err := client.CreateExec(docker.CreateExecOptions{
		AttachStdin:  stdin,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          command,
		Container:    container,
		User:         "root",
	})
	if err != nil {
		fmt.Println("Error creating execution:", err)
		return
	}

	// start the command that is ready to run
	err = client.StartExec(exec.ID, docker.StartExecOptions{
		Detach:       false,
		Tty:          true,
		InputStream:  os.Stdin,
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
		RawTerminal:  true,
	})
	if err != nil {
		fmt.Println("Error starting execution:", err)
		return
	}
}

// RunLinuxCmd just does a normal linux command. This is used if we can't
// run a certain command, or if a command was run out of the context of the /containers
// directory.
func RunLinuxCmd(c *cli.Context) {
	//fmt.Println("Not container directory")
	env := os.Getenv("PATH")
	split := strings.Split(env, ":")
	newEnv := make([]string, len(split)-1)
	for _, entry := range split {
		if entry != "/kodybin" {
			newEnv = append(newEnv, entry)
		}
	}
	env = strings.Join(newEnv, ":")
	os.Setenv("PATH", env)
	//fmt.Println("PATH is:", os.Getenv("PATH"))

	cmd := exec.Command(c.App.Name, c.Args()...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
	fmt.Println(out.String())
}

// GetContainerName takes the string the user provided on the command line
// and extracts the container name.
// Some example inputs: '/containers/awesome_cray' or 'awesome_cray/var/lib'
func GetContainerName(arg string) (string, error) {
	if arg == "" {
		return "", fmt.Errorf("Empty string provided")
	}

	split := strings.Split(arg, "/")
	if split[0] == "" {
		split = split[1:]
	}

	//fmt.Printf("Decoded %q\n", split)
	if len(split) == 1 && split[0] == "containers" {
		//fmt.Println("Returning \"\"")
		return "", nil // user provided '/containers'
	}
	if len(split) > 1 && strings.Index(split[1], "swarm-") == 0 {
		//fmt.Println("Swarm node included in name")
		return strings.Join(split[1:], "/"), nil //swarm support
	}
	if len(split) > 1 && split[0] == "containers" {
		//fmt.Printf("Returning %q\n", split[1])
		return split[1], nil // user provided '/containers/<container name>'
	}
	if len(split) > 1 && split[0] != "containers" {
		//fmt.Println("got containername/subdir")
		return split[0], nil // user provided '<name>/<subdir>'
	}
	fmt.Println("Just container name")
	return arg, nil // user provided just a container name
}

// IsContainerDirectory checks to see if the current directory is the /containers
// directory. If it is, then all of terminal commands should turn into terminal commands.
func IsContainerDirectory(c *cli.Context) bool {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return false
	}

	if dir == "/containers" {
		return true
	}
	return false
}
