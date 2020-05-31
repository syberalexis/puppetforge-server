package controller

import (
	"net/http"

	"github.com/syberalexis/puppetforge-server/pkg/bridge"
	"github.com/syberalexis/puppetforge-server/pkg/forge"
)

type ControllerContext struct {
	Bridge         *bridge.ForgeModuleBridge
	ModuleService  *forge.ModuleService
	ReleaseService *forge.ReleaseService
}

type ControllerHandler struct {
	*ControllerContext
	//ContextedHandlerFunc is the interface which our Handlers will implement
	ContextedHandlerFunc func(*ControllerContext, http.ResponseWriter, *http.Request)
}

func (handler ControllerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.ContextedHandlerFunc(handler.ControllerContext, w, r)
}
