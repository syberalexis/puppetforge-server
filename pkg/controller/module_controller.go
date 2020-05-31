package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syberalexis/puppetforge-server/pkg/model"
)

func FetchModule(context *ControllerContext, w http.ResponseWriter, r *http.Request) {
	prepareResponse(w)
	slugModule := extractVariables(r)

	module, err := context.ModuleService.GetModule(slugModule)
	if err != nil {
		module, err = context.Bridge.FetchModule(slugModule)
	}

	if err != nil {
		json.NewEncoder(w).Encode(model.ForgeError{Message: "500 : Bridge error", Errors: []string{err.Error()}})
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(module)
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteModule(context *ControllerContext, w http.ResponseWriter, r *http.Request) {

}

func DeprecateModule(context *ControllerContext, w http.ResponseWriter, r *http.Request) {

}

func ListModules(context *ControllerContext, w http.ResponseWriter, r *http.Request) {
	prepareResponse(w)

	module, err := context.Bridge.ListModules(0, 0, "maeq-thanos", "", "", "", false, false, false, []string{}, "", "", "", 0, []string{}, false, false, false, []string{}, false, []string{}, []string{}, false)
	if err != nil {
		json.NewEncoder(w).Encode(model.ForgeError{Message: "500 : Bridge error", Errors: []string{err.Error()}})
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(module)
		w.WriteHeader(http.StatusOK)
	}
}

func prepareResponse(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
}

func extractVariables(r *http.Request) (slugModule string) {
	vars := mux.Vars(r)
	slugModule = vars["module_slug"]
	// strconv.Atoi is shorthand for ParseInt
	//id, err := strconv.Atoi(vars["id"])
	return
}
