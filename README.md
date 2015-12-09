# dockerbox
We all know and love busybox, right? busybox is a single compiled
binary that intercepts multiple common terminal commands.

We all know and love Docker, right? Docker gives devops engineers the
ability to build, ship, and run applications in a packaged environment.

What if we put the two together and gave devops engineers a command
line interface that they can use to gain control over their Docker
deployments?

## Terminal commands we know and love
Docker is famous for using a combo of Linux/git primitives in their
command line interface. That's great, but as time has progressed,
the Docker command line has grown quite large and confusing. Let's
bring the Docker cli back to its roots - Linux!

With dockerbox, we intercept common Linux terminal commands, translate
them into their Docker cli counterparts, and then display the results
in a way that a normal Linux user would expect.

For example, take the `ls` command. Normally, an `ls` will list all
files and subdirectories of a given directory.

Here's what `ls` does on my system.

```
$ ls /
Applications              cores                     opt
Library                   dev                       private
Network                   etc                       sbin
System                    home                      tmp
Users                     installer.failurerequests usr
Volumes                   kodybin                   var
bin                       net
```

Hm. That's cool. Now let's see how we can `ls /` inside of a
container.

```
# first we need to know the name of the container
$ docker ps
CONTAINER ID        IMAGE               COMMAND               CREATED             STATUS              PORTS               NAMES
06b6b108779b        google/cadvisor     "/usr/bin/cadvisor"   2 seconds ago       Up 1 seconds        8080/tcp            swarm-agent-01/cadvisor

# now that we have the name (at the end), we can run
# another command to get a shell
$ docker exec -it swarm-agent-01/cadvisor /bin/sh
/ #
```

That seems like a lot of work just to get a shell in a container.

What if we could just use the commands we already know? That's what
dockerbox does! Once it's installed, this is how we do the same thing.

```
# find the name of the container by listing containers
$ ls /containers
ls /containers

swarm-agent-01/cadvisor             swarm-agent-00/swarm-agent     swarm-master/swarm-agent
swarm-master/swarm-agent-master     swarm-agent-01/swarm-agent

# now that we have the name, let's 'cd' into it
$ cc /containers/swarm-agent-01/cadvisor
/ #
```

## So there we go
No longer do we have to watch GitHub issues and pull requests just to
keep up with the changing Docker cli. We can just use the one we
already know!

## Some added goodies
The `ls /containers` command will colorize the output! If a container
is started, it will appear green. If a container is stopped, it
appears red.

A special `stop` and `start` command are included to start and stop
containers.

## Future
In an ideal world, this would completely replace the Docker cli.
For 80% of use cases, it's probably fine that you don't get all of
the nitty gritty details and flags that the Docker cli supports.

Future tasks:

* Separate listed containers into host-based subdirectories
* Have the `cc` command take a path to navigate to inside the container
* (big one) Make this binary into an interactive shell
  * This would allow us to not worry about messing up terminal
  commands
  * Plus, we could actually use the `cd` command instead of `cc`
* Volume support
  * Create volumes using terminal commands
  * Specify the volume driver by using a subdirectory
  * Imagine `ls /volumes` and `mkdir /volumes/veritas/myVol`
  * Imagine `ln -s /volumes/veritas/myVol
  /containers/swarm-01/mysql/data`
