package docker

import (
	"github.com/docker/docker/client"
)

// Docker client API version
var version = "1.40"

// getDocker return Docker client.
func getDocker() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion(version))
	if err != nil {
		panic(err)
	}
	return cli
}
