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
func InfoDocker() (types.Info, error) {
	info, err := docker.Info(context.Background())
	if err != nil {
		return types.Info{}, err
	}
	return info, nil
}

/** --- Container --- */

// ListContainers return containers.
func ListContainers() ([]types.Container, error) {
	containers, err := docker.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

// StartContainer start container.
func StartContainer(containerID string) error {
	err := docker.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

// StopContainer stop container.
func StopContainer(containerID string) error {
	err := docker.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		return err
	}
	return nil
}

// RemoveContainer remove container.
func RemoveContainer(containerID string) error {
	err := docker.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

// InspectContainer inspect specific container.
func InspectContainer(containerID string) (types.ContainerJSON, error) {
	inspect, err := docker.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return inspect, nil
}

// CommitContainer commit container.
func CommitContainer(containerID string, repo string, tag string) error {
	_, err := docker.ContainerCommit(context.Background(), containerID, types.ContainerCommitOptions{
		Reference: repo + ":" + tag,
	})
	if err != nil {
		return err
	}
	return nil
}

// PauseContainer pause container.
func PauseContainer(containerID string) error {
	err := docker.ContainerPause(context.Background(), containerID)
	if err != nil {
		return err
	}
	return nil
}

// UnpauseContainer unpause container.
func UnpauseContainer(containerID string) error {
	err := docker.ContainerUnpause(context.Background(), containerID)
	if err != nil {
		return err
	}
	return nil
}

// RenameContainer rename container.
func RenameContainer(containerID string, newContainerName string) error {
	err := docker.ContainerRename(context.Background(), containerID, newContainerName)
	if err != nil {
		return err
	}
	return nil
}

/** --- Image --- */

// ListImages return images.
func ListImages() ([]types.ImageSummary, error) {
	images, err := docker.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}
	return images, nil
}

// RemoveImage remove image.
func RemoveImage(imageID string) error {
	_, err := docker.ImageRemove(context.Background(), imageID, types.ImageRemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

// SearchImage search images from DockerHub.
func SearchImage(keyword string) ([]registry.SearchResult, error) {
	images, err := docker.ImageSearch(context.Background(), keyword, types.ImageSearchOptions{Limit: 100})
	if err != nil {
		return nil, err
	}
	return images, nil
}

// PullImage pull image.
func PullImage(name string) error {
	closer, err := docker.ImagePull(context.Background(), name, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer closer.Close()

	scanner := bufio.NewScanner(closer)
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
