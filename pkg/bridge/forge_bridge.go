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
	Uri string
}

// FetchModule method to fetch a module by his name from a Puppet forge server
func (forge *ForgeBridge) FetchModule(name string) (module model.Module, err error) {
	err = executeRequest(fmt.Sprintf("%s/v3/modules/%s", forge.Uri, name), &module)
	return
}

// ListModules method to list modules with search query
func (forge *ForgeBridge) ListModules(limit uint, offset uint, sortBy string, query string, tag string, owner string, withTasks bool,
	withPlans bool, withPdk bool, endorsements []string, operatingsystem string, peRequirement string, puppetRequirement string,
	withMinimumScore int, moduleGroups []string, showDeleted bool, hideDeprecated bool, onlyLatest bool, slugs []string, withHtml bool,
	includeFields []string, excludeFields []string, supported bool) (modules []model.Module, err error) {

	var page model.Page
	var queryParams string
	params := make([]string, 0)

	// Format query parameters
	if limit > 0 {
		params = append(params, fmt.Sprintf("%s=%s", "limit", string(limit)))
	}
	if offset > 0 {
		params = append(params, fmt.Sprintf("%s=%s", "offset", string(offset)))
	}
	if sortBy != "" {
		params = append(params, fmt.Sprintf("%s=%s", "sort_by", sortBy))
	}
	if query != "" {
		params = append(params, fmt.Sprintf("%s=%s", "query", query))
	}
	if tag != "" {
		params = append(params, fmt.Sprintf("%s=%s", "tag", tag))
	}
	if owner != "" {
		params = append(params, fmt.Sprintf("%s=%s", "owner", owner))
	}
	if withTasks {
		params = append(params, fmt.Sprintf("%s=%s", "with_tasks", "true"))
	}
	if withPlans {
		params = append(params, fmt.Sprintf("%s=%s", "with_plans", "true"))
	}
	if withPdk {
		params = append(params, fmt.Sprintf("%s=%s", "with_pdk", "true"))
	}
	if len(endorsements) > 0 {
		params = append(params, fmt.Sprintf("%s=%s", "endorsements", strings.Join(endorsements, ",")))
	}
	if operatingsystem != "" {
		params = append(params, fmt.Sprintf("%s=%s", "operatingsystem", operatingsystem))
	}
	if peRequirement != "" {
		params = append(params, fmt.Sprintf("%s=%s", "pe_requirement", peRequirement))
	}
	if puppetRequirement != "" {
		params = append(params, fmt.Sprintf("%s=%s", "puppet_requirement", puppetRequirement))
	}
	if withMinimumScore > 0 {
		params = append(params, fmt.Sprintf("%s=%s", "with_minimum_score", string(withMinimumScore)))
	}
	if len(moduleGroups) > 0 {
		params = append(params, fmt.Sprintf("%s=%s", "module_groups", strings.Join(moduleGroups, ",")))
	}
	if showDeleted {
		params = append(params, fmt.Sprintf("%s=%s", "show_deleted", "true"))
	}
	if hideDeprecated {
		params = append(params, fmt.Sprintf("%s=%s", "hide_deprecated", "true"))
	}
	if onlyLatest {
		params = append(params, fmt.Sprintf("%s=%s", "only_latest", "true"))
	}
	if len(slugs) > 0 {
		params = append(params, fmt.Sprintf("%s=%s", "slugs", strings.Join(slugs, ",")))
	}
	if withHtml {
		params = append(params, fmt.Sprintf("%s=%s", "with_html", "true"))
	}
	if len(includeFields) > 0 {
		params = append(params, fmt.Sprintf("%s=%s", "include_fields", strings.Join(includeFields, ",")))
	}
	if len(excludeFields) > 0 {
		params = append(params, fmt.Sprintf("%s=%s", "exclude_fields", strings.Join(excludeFields, ",")))
	}
	if supported {
		params = append(params, fmt.Sprintf("%s=%s", "supported", "true"))
	}

	if len(params) > 0 {
		queryParams = fmt.Sprintf("?%s", strings.Join(params, "&"))
	} else {
		queryParams = ""
	}

	err = executeRequest(fmt.Sprintf("%s/v3/modules%s", forge.Uri, queryParams), &page)
	modules = page.Results
	return
}

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
