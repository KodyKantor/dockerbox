package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

// RunCmd takes the name of a container, and a command to run (a list of arguments).
// It sets up a 'docker exec'-type command and executes it.
func RunCmd(container string, command []string, detach bool) {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		fmt.Println("Couldn't connect to docker:", err)
	}

	exec, err := client.CreateExec(docker.CreateExecOptions{
		AttachStdin:  !detach,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          command,
		Container:    container,
		User:         "root",
	})
	if err != nil {
		fmt.Println("Error creating execution:", err)
	}

	err = client.StartExec(exec.ID, docker.StartExecOptions{
		Detach:       !detach,
		Tty:          true,
		InputStream:  os.Stdin,
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
		RawTerminal:  true,
	})
	if err != nil {
		fmt.Println("Error starting execution:", err)
	}
}

// GetContainerName takes the string the user provided on the command line
// and extracts the container name.
// Some example inputs: '/containers/awesome_cray' or 'awesome_cray/var/lib'
func GetContainerName(arg string) (string, error) {
	if arg == "" {
		return "", fmt.Errorf("Empty string provided")
	}

	split := strings.Split(arg, "/")
	//fmt.Printf("Decoded %q\n", split)
	if split[0] == "" {
		split = split[1:]
	}

	if len(split) == 1 && split[0] == "containers" {
		fmt.Println("Returning \"\"")
		return "", nil // user provided '/containers'
	}
	if len(split) == 2 && split[0] == "containers" {
		fmt.Printf("Returning %q\n", split[1])
		return split[1], nil // user provided '/containers/<container name>'
	}
	if len(split) > 1 && split[0] != "containers" {
		return split[0], nil // user provided '<name>/<subdir>'
	}
	return arg, nil // user provided just a container name
}
