package appspec

// App represents the appspec, which is a collection of Processes. Each process
// being a collection of containers.
type App struct {
	// The name of this app.
	Name string

	Processes []Process
}

type Process struct {
	// The type of this processes (e.g. "web").
	Type string

	// The number of instances of this process to run.
	Instances int

	// An array of containers that make up this process.
	Containers []Container
}

type Container struct {
	// The name of this container. This can be used to link
	// to another container.
	Name string

	// The container image to run.
	Image string

	// The command to run.
	Command []string

	// Environment variables to set.
	Env map[string]string

	// Memory limit.
	Memory int64

	// CPU shares to limit.
	CPU int64

	// Links to other containers within the same Process.
	Links []string

	// Ports to allocate to this container.
	Ports []Port
}

type Port struct {
	// Controls whether this port should have a load balancer attached to
	// it.
	Exposure string

	// The port exposed on the host.
	Host *int64

	// The port that the process within the container binds to.
	Container *int64

	// The protocol. Can be "tcp" or "udp".
	Protocol *string
}
