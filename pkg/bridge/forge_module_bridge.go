package bridge

import (
	"fmt"
	"strings"

	"github.com/syberalexis/puppetforge-server/pkg/model"
)

// ForgeModuleBridge Bridge Puppet Forge's modules endpoints
type ForgeModuleBridge struct {
	URI string
}

// FetchModule Returns data for a single Module resource identified by the module's slug value.
func (forge *ForgeModuleBridge) FetchModule(name string) (module *model.Module, err error) {
	err = executeRequest(fmt.Sprintf("%s/v3/modules/%s", forge.URI, name), module)
	return
}

// ListModules Returns a list of modules meeting the specified search criteria and filters. Results are paginated.
func (forge *ForgeModuleBridge) ListModules(limit uint, offset uint, sortBy string, query string, tag string, owner string, withTasks bool,
	withPlans bool, withPdk bool, endorsements []string, operatingsystem string, peRequirement string, puppetRequirement string,
	withMinimumScore int, moduleGroups []string, showDeleted bool, hideDeprecated bool, onlyLatest bool, slugs []string, withHTML bool,
	includeFields []string, excludeFields []string, supported bool) (modules []model.Module, err error) {

	var page model.Page
	var queryParams string
	params := make([]string, 0)

	// Format query parameters
	if limit > 0 {
		params = append(params, fmt.Sprintf("%s=%d", "limit", limit))
	}
	if offset > 0 {
		params = append(params, fmt.Sprintf("%s=%d", "offset", offset))
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
		params = append(params, fmt.Sprintf("%s=%d", "with_minimum_score", withMinimumScore))
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
	if withHTML {
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

	err = executeRequest(fmt.Sprintf("%s/v3/modules%s", forge.URI, queryParams), &page)
	modules = page.Results
	return
}
