package server

import (
	"github.com/gorilla/mux"
	"github.com/syberalexis/puppetforge-server/pkg/controller"
	"github.com/syberalexis/puppetforge-server/pkg/spa"
)

// NewRouter construct mux router to expose endpoints and single app
func NewRouter(context *controller.ControllerContext) *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	// React App
	router.PathPrefix("/").Handler(&spa.SpaHandler{StaticPath: "web/build", IndexPath: "index.html"})
	// Module endpoints
	router.Methods("GET").Path("/v3/modules/{module_slug}").Name("Get Module").Handler(&controller.ControllerHandler{context, controller.FetchModule})
	router.Methods("DELETE").Path("/v3/modules/{module_slug}").Name("Delete Module").Handler(&controller.ControllerHandler{context, controller.DeleteModule})
	router.Methods("PATCH").Path("/v3/modules/{module_slug}").Name("Delete Module").Handler(&controller.ControllerHandler{context, controller.DeprecateModule})
	router.Methods("GET").Path("/v3/modules").Name("List Modules").Handler(&controller.ControllerHandler{context, controller.ListModules})
	// Release endpoints
	router.Methods("GET").Path("/v3/releases").Name("List module releases").Handler(&controller.ControllerHandler{context, controller.ListReleases})
	router.Methods("POST").Path("/v3/releases").Name("Create module release").Handler(&controller.ControllerHandler{context, controller.CreateRelease})
	router.Methods("GET").Path("/v3/releases/{release_slug}").Name("Fetch module release").Handler(&controller.ControllerHandler{context, controller.FetchRelease})
	router.Methods("DELETE").Path("/v3/releases/{release_slug}").Name("Delete module release").Handler(&controller.ControllerHandler{context, controller.DeleteRelease})
	router.Methods("GET").Path("/v3/releases/{release_slug}/plans").Name("List module release plans").Handler(&controller.ControllerHandler{context, controller.ListReleasePlans})
	router.Methods("GET").Path("/v3/releases/{release_slug}/plans/{plan_name}").Name("Fetch module release plan").Handler(&controller.ControllerHandler{context, controller.FetchReleasePlan})
	router.Methods("GET").Path("/v3/files/{filename}").Name("Download module release").Handler(&controller.ControllerHandler{context, controller.DownloadRelease})

	return router
}
