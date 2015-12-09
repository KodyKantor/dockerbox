package cd

import "github.com/codegangsta/cli"

// Chdir implements the Operationer interface.
type Chdir struct{}

// DoStuff will create a shell session inside the chosen container.
func (cd *Chdir) DoStuff(c *cli.Context) {

}
