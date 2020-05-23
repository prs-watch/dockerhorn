package main

import (
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/prs-watch/dockerhorn/docker"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(&rest.CorsMiddleware{
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
	})
	router, err := rest.MakeRouter(
		// routes for docker
		rest.Get("/info", infoDocker),
		// routes for container
		rest.Get("/container/ps", listContainers),
		rest.Get("/container/:containerID", inspectContainer),
		rest.Get("/container/start/:containerID", startContainer),
		rest.Get("/container/stop/:containerID", stopContainer),
		rest.Get("/container/remove/:containerID", removeContainer),
		rest.Get("/container/commit/:containerID", commitContainer),
		rest.Get("/container/pause/:containerID", pauseContainer),
		rest.Get("/container/unpause/:containerID", unpauseContainer),
		rest.Get("/container/rename/:containerID", renameContainer),
		// routes for image
		rest.Get("/image/images", listImages),
		rest.Get("/image/remove/:imageID", removeImage),
		rest.Get("/image/search/:keyword", searchImage),
		rest.Get("/image/pull/:name", pullImage),
	)
	if err != nil {
		panic(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9999", api.MakeHandler()))
}

/** --- Docker ---*/

func infoDocker(w rest.ResponseWriter, r *rest.Request) {
	info, err := docker.InfoDocker()
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]types.Info{
		"Info": info,
	})
}

/** --- Container --- */

// listContainers write containers' info.
func listContainers(w rest.ResponseWriter, r *rest.Request) {
	containers, err := docker.ListContainers()
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string][]types.Container{
		"Containers": containers,
	})
}

// inspectContainer write container inspected info.
func inspectContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	inspect, err := docker.InspectContainer(containerID)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]types.ContainerJSON{
		"Inspect": inspect,
	})
}

// startContainer start container.
func startContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	err := docker.StartContainer(containerID)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// stopContainer stop container.
func stopContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	err := docker.StopContainer(containerID)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// removeContainer remove container.
func removeContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	err := docker.RemoveContainer(containerID)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// commitContainer commit container.
func commitContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	query := r.URL.Query()
	repo := query["repo"][0]
	tag := query["tag"][0]
	err := docker.CommitContainer(containerID, repo, tag)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// pauseContainer pause container.
func pauseContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	err := docker.PauseContainer(containerID)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// pauseContainer pause container.
func unpauseContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	err := docker.UnpauseContainer(containerID)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// renameContainer rename container.
func renameContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	query := r.URL.Query()
	name := query["name"][0]
	err := docker.RenameContainer(containerID, name)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

/** --- Image --- */

// listImages write images' info.
func listImages(w rest.ResponseWriter, r *rest.Request) {
	images, err := docker.ListImages()
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string][]types.ImageSummary{
		"Images": images,
	})
}

// removeImage remove image.
func removeImage(w rest.ResponseWriter, r *rest.Request) {
	imageID := r.PathParam("imageID")
	err := docker.RemoveImage(imageID)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// searchImage search images from DockerHub.
func searchImage(w rest.ResponseWriter, r *rest.Request) {
	keyword := r.PathParam("keyword")
	images, err := docker.SearchImage(keyword)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string][]registry.SearchResult{
		"Images": images,
	})
}

// pullImage pull image.
func pullImage(w rest.ResponseWriter, r *rest.Request) {
	name := r.PathParam("name")
	err := docker.PullImage(name)
	if err != nil {
		w.WriteJson(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}
