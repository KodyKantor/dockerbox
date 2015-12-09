package operation

import (
	"fmt"

	"github.com/codegangsta/cli"
	"stash.veritas.com/scm/kody/dockerbox/operation/cd"
	"stash.veritas.com/scm/kody/dockerbox/operation/fallback"
	"stash.veritas.com/scm/kody/dockerbox/operation/ls"
)

//Operationer interface is an abstraction for terminal operations.
type Operationer interface {
	DoStuff(*cli.Context)
}

//GetOperation returns an instance of an Operation interface. The caller
// is then intended to run DoStuff on that returned Operation.
func GetOperation(name string) Operationer {
	switch name {
	case "ls":
		return &ls.List{}
	case "cc":
		return &cd.Chdir{}
	default:
		fmt.Println("Falling back!")
		return &fallback.Fallback{Operation: name}
	}
}
