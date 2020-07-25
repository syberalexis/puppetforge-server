package server

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/syberalexis/puppetforge-server/pkg/bridge"
	"github.com/syberalexis/puppetforge-server/pkg/controller"
	"github.com/syberalexis/puppetforge-server/pkg/forge"
)

// PuppetForge Object to run and serve The Private Puppet Forge Server
type PuppetForge struct {
	Host     string
	Port     uint
	ForgeURI string
}

// Run Method to run the Private Puppet Forge Server
func (puppetforge *PuppetForge) Run() {
	router := NewRouter(&controller.ControllerContext{
		Bridge:         &bridge.ForgeModuleBridge{URI: puppetforge.ForgeURI},
		ModuleService:  &forge.ModuleService{},
		ReleaseService: &forge.ReleaseService{},
	})

	log.Infof("Server is listen on %s:%d", puppetforge.Host, puppetforge.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", puppetforge.Host, puppetforge.Port), router)
}
