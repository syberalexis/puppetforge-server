package server

import (
	"fmt"
	"net/http"

	"github.com/syberalexis/puppetforge-server/pkg/bridge"
	"github.com/syberalexis/puppetforge-server/pkg/controller"
)

type PuppetForge struct {
	Host     string
	Port     uint
	ForgeURI string
}

func (puppetforge *PuppetForge) Run() {
	router := NewRouter(&controller.ControllerContext{
		Bridge: &bridge.ForgeModuleBridge{URI: puppetforge.ForgeURI},
	})

	http.ListenAndServe(fmt.Sprintf("%s:%d", puppetforge.Host, puppetforge.Port), router)
}
