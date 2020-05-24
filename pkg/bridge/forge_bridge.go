package bridge

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/syberalexis/puppetforge-server/pkg/model"
)

type ForgeBridge struct {
	uri string
}

func (forge *ForgeBridge) FetchModule(name string) (module model.Module, err error) {
	var forgeError model.ForgeError

	// Get module from Forge API
	resp, err := http.Get(fmt.Sprintf("%s/v3/modules/%s", forge.uri, name))
	if err != nil {
		return
	}

	// Read Body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Unmarshal module or error
	if resp.StatusCode != 200 {
		err = json.Unmarshal(body, &forgeError)
		if err != nil {
			return
		}
		err = errors.New(fmt.Sprintf("%s :\n%s", forgeError.Message, strings.Join(forgeError.Errors, "\n")))
		return
	} else {
		err = json.Unmarshal(body, &module)
		return
	}
}
