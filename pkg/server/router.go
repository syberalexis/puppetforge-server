package server

import (
	"github.com/gorilla/mux"
	"github.com/syberalexis/puppetforge-server/pkg/controller"
)

func NewRouter(context *controller.ControllerContext) *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/v3/modules/{module_slug}").Name("Get Module").Handler(&controller.ControllerHandler{context, controller.FetchModule})
	router.Methods("DELETE").Path("/v3/modules/{module_slug}").Name("Delete Module").Handler(&controller.ControllerHandler{context, controller.DeleteModule})
	router.Methods("PATCH").Path("/v3/modules/{module_slug}").Name("Delete Module").Handler(&controller.ControllerHandler{context, controller.DeprecateModule})
	router.Methods("GET").Path("/v3/modules").Name("List Modules").Handler(&controller.ControllerHandler{context, controller.ListModules})
	return router
}
