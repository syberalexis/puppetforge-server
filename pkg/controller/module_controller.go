package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/syberalexis/puppetforge-server/pkg/model"
)

// FetchModule Method to fetch a Puppet module from local storage or another Forge
func FetchModule(context *ControllerContext, w http.ResponseWriter, r *http.Request) {
	prepareResponse(w)
	vars := mux.Vars(r)
	slugModule := vars["module_slug"]

	module, err := context.ModuleService.GetModule(slugModule)
	if module == nil {
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

// DeleteModule Method to delete a local Puppet module
func DeleteModule(context *ControllerContext, w http.ResponseWriter, r *http.Request) {

}

// DeprecateModule Method to deprecate a local Puppet module
func DeprecateModule(context *ControllerContext, w http.ResponseWriter, r *http.Request) {

}

// ListModules Method to search and list Puppet modules form local storage or another Forge
func ListModules(context *ControllerContext, w http.ResponseWriter, r *http.Request) {
	prepareResponse(w)

	limit, err := strconv.ParseUint(r.FormValue("limit"), 10, 64)
	if err != nil {
		limit = 0
	}
	offset, err := strconv.ParseUint(r.FormValue("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	sortBy := r.FormValue("sortBy")
	query := r.FormValue("query")
	tag := r.FormValue("tag")
	owner := r.FormValue("owner")
	withTasks, err := strconv.ParseBool(r.FormValue("withTasks"))
	if err != nil {
		withTasks = false
	}
	withPlans, err := strconv.ParseBool(r.FormValue("lwithPlansimit"))
	if err != nil {
		withPlans = false
	}
	withPdk, err := strconv.ParseBool(r.FormValue("withPdk"))
	if err != nil {
		withPdk = false
	}
	endorsements := parseStringArray(r.FormValue("endorsements"))
	operatingsystem := r.FormValue("operatingsystem")
	peRequirement := r.FormValue("peRequirement")
	puppetRequirement := r.FormValue("puppetRequirement")
	withMinimumScore, err := strconv.ParseInt(r.FormValue("withMinimumScore"), 10, 64)
	if err != nil {
		withMinimumScore = 0
	}
	moduleGroups := parseStringArray(r.FormValue("moduleGroups"))
	showDeleted, err := strconv.ParseBool(r.FormValue("showDeleted"))
	if err != nil {
		showDeleted = false
	}
	hideDeprecated, err := strconv.ParseBool(r.FormValue("hideDeprecated"))
	if err != nil {
		hideDeprecated = false
	}
	onlyLatest, err := strconv.ParseBool(r.FormValue("onlyLatest"))
	if err != nil {
		onlyLatest = false
	}
	slugs := parseStringArray(r.FormValue("slugs"))
	withHTML, err := strconv.ParseBool(r.FormValue("withHTML"))
	if err != nil {
		withHTML = false
	}
	includeFields := parseStringArray(r.FormValue("includeFields"))
	excludeFields := parseStringArray(r.FormValue("excludeFields"))
	supported, err := strconv.ParseBool(r.FormValue("supported"))
	if err != nil {
		supported = false
	}

	// TODO Search from local database

	module, err := context.Bridge.ListModules(uint(limit), uint(offset), sortBy, query, tag, owner, withTasks, withPlans, withPdk, endorsements, operatingsystem,
		peRequirement, puppetRequirement, int(withMinimumScore), moduleGroups, showDeleted, hideDeprecated, onlyLatest,
		slugs, withHTML, includeFields, excludeFields, supported)
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

func parseStringArray(stringArray string) (res []string) {
	if len(stringArray) == 0 {
		res = []string{}
	} else {
		res = strings.Split(stringArray, ",")
	}
	return
}
