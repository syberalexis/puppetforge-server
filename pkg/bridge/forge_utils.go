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

func executeRequest(uri string, result interface{}) error {
	var forgeError model.ForgeError

	// Execute request from Forge API
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	// Read Body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal module or error
	if resp.StatusCode != 200 {
		err = json.Unmarshal(body, &forgeError)
		if err != nil {
			return err
		}
		err = errors.New(fmt.Sprintf("%s :\n%s", forgeError.Message, strings.Join(forgeError.Errors, "\n")))
		return err
	} else {
		err = json.Unmarshal(body, &result)
		return err
	}
}
