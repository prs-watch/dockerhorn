package docker

import (
	"bufio"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
)

// docker Docker client
var docker = getDocker()

/** --- Docker --- */

// InfoDocker return docker info.
func InfoDocker() types.Info {
	info, err := docker.Info(context.Background())
	if err != nil {
		panic(err)
	}
	return info
}

/** --- Container --- */

// ListContainers return containers.
func ListContainers() []types.Container {
	containers, err := docker.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}
	return containers
}

// StartContainer start container.
func StartContainer(containerID string) {
	err := docker.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
}

// StopContainer stop container.
func StopContainer(containerID string) {
	err := docker.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		panic(err)
	}
}

// RemoveContainer remove container.
func RemoveContainer(containerID string) {
	err := docker.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{})
	if err != nil {
		panic(err)
	}
}

// InspectContainer inspect specific container.
func InspectContainer(containerID string) types.ContainerJSON {
	inspect, err := docker.ContainerInspect(context.Background(), containerID)
	if err != nil {
		panic(err)
	}
	return inspect
}

// CommitContainer commit container.
func CommitContainer(containerID string, repo string, tag string) {
	_, err := docker.ContainerCommit(context.Background(), containerID, types.ContainerCommitOptions{
		Reference: repo + ":" + tag,
	})
	if err != nil {
		panic(err)
	}
}

/** --- Image --- */

// ListImages return images.
func ListImages() []types.ImageSummary {
	images, err := docker.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	return images
}

// RemoveImage remove image.
func RemoveImage(imageID string) {
	_, err := docker.ImageRemove(context.Background(), imageID, types.ImageRemoveOptions{})
	if err != nil {
		panic(err)
	}
}

// SearchImage search images from DockerHub.
func SearchImage(keyword string) []registry.SearchResult {
	images, err := docker.ImageSearch(context.Background(), keyword, types.ImageSearchOptions{Limit: 100})
	if err != nil {
		panic(err)
	}
	return images
}

// PullImage pull image.
func PullImage(name string) {
	closer, err := docker.ImagePull(context.Background(), name, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	scanner := bufio.NewScanner(closer)
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
