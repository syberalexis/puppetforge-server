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
	ForgeUri string
}

func (puppetforge *PuppetForge) Run() {
	router := NewRouter(&controller.ControllerContext{
		Bridge: &bridge.ForgeBridge{Uri: puppetforge.ForgeUri},
	})

	http.ListenAndServe(fmt.Sprintf("%s:%d", puppetforge.Host, puppetforge.Port), router)
}

func TestModel() {
	bridge := bridge.ForgeBridge{Uri: "https://forgeapi.puppet.com"}
	modules, err := bridge.ListModules(0, 0, "", "maeq", "", "", false, false, false, []string{}, "", "", "", 0, []string{}, false, false, false, []string{}, false, []string{}, []string{}, false)

	fmt.Println(err)
	for _, module := range modules {
		fmt.Println(module.SlugName)
	}
}
