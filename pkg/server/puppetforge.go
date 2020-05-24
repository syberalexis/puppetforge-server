package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/syberalexis/puppetforge-server/pkg/model"
)

func TestModel() {
	var module model.Module
	var forgeError model.ForgeError
	resp, _ := http.Get("https://forgeapi.puppet.com/v3/modules/maeq-thanos2")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 204 {
		json.Unmarshal(body, &forgeError)
	} else {
		json.Unmarshal(body, &module)
	}
	fmt.Println(forgeError)
	fmt.Println(module.Owner.Uri)
	fmt.Println(module.Owner.SlugName)
	fmt.Println(module.Owner.Username)
	fmt.Println(module.Owner.GravatarId)
}
