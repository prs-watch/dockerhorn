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
	info := docker.InfoDocker()
	w.WriteJson(map[string]types.Info{
		"Info": info,
	})
}

/** --- Container --- */

// listContainers write containers' info.
func listContainers(w rest.ResponseWriter, r *rest.Request) {
	containers := docker.ListContainers()
	w.WriteJson(map[string][]types.Container{
		"Containers": containers,
	})
}

// inspectContainer write container inspected info.
func inspectContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	inspect := docker.InspectContainer(containerID)
	w.WriteJson(map[string]types.ContainerJSON{
		"Inspect": inspect,
	})
}

// startContainer start container.
func startContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	docker.StartContainer(containerID)
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// stopContainer stop container.
func stopContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	docker.StopContainer(containerID)
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// removeContainer remove container.
func removeContainer(w rest.ResponseWriter, r *rest.Request) {
	containerID := r.PathParam("containerID")
	docker.RemoveContainer(containerID)
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
	docker.CommitContainer(containerID, repo, tag)
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

/** --- Image --- */

// listImages write images' info.
func listImages(w rest.ResponseWriter, r *rest.Request) {
	images := docker.ListImages()
	w.WriteJson(map[string][]types.ImageSummary{
		"Images": images,
	})
}

// removeImage remove image.
func removeImage(w rest.ResponseWriter, r *rest.Request) {
	imageID := r.PathParam("imageID")
	docker.RemoveImage(imageID)
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}

// searchImage search images from DockerHub.
func searchImage(w rest.ResponseWriter, r *rest.Request) {
	keyword := r.PathParam("keyword")
	images := docker.SearchImage(keyword)
	w.WriteJson(map[string][]registry.SearchResult{
		"Images": images,
	})
}

// pullImage pull image.
func pullImage(w rest.ResponseWriter, r *rest.Request) {
	name := r.PathParam("name")
	docker.PullImage(name)
	w.WriteJson(map[string]string{
		"Status": "OK",
	})
}
