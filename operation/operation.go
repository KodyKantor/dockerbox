package operation

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/kodykantor/dockerbox/operation/cd"
	"github.com/kodykantor/dockerbox/operation/fallback"
	"github.com/kodykantor/dockerbox/operation/ls"
	"github.com/kodykantor/dockerbox/operation/start"
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
	case "start":
		return &start.Start{}
	default:
		fmt.Println("Falling back!")
		return &fallback.Fallback{Operation: name}
	}
}
