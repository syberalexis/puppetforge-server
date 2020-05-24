package server

import (
	"fmt"

	"github.com/syberalexis/puppetforge-server/pkg/bridge"
)

func TestModel() {
	bridge := bridge.ForgeBridge{Uri: "https://forgeapi.puppet.com"}
	modules, err := bridge.ListModules(0, 0, "", "maeq", "", "", false, false, false, []string{}, "", "", "", 0, []string{}, false, false, false, []string{}, false, []string{}, []string{}, false)

	fmt.Println(err)
	for _, module := range modules {
		fmt.Println(module.SlugName)
	}
}
