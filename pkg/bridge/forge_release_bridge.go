package bridge

import (
	"fmt"

	"github.com/syberalexis/puppetforge-server/pkg/model"
)

// ForgeReleaseBridge Bridge Puppet Forge's releases endpoints
type ForgeReleaseBridge struct {
	URI string
}

// ListReleases Returns a list of module releases meeting the specified search criteria and filters. Results are paginated.
func (forge *ForgeReleaseBridge) ListReleases(name string) (module model.Module, err error) {
	err = executeRequest(fmt.Sprintf("%s/v3/releases", forge.URI), &module)
	return
}

// FetchRelease Returns data for a single module Release resource identified by the module release's slug value.
func (forge *ForgeReleaseBridge) FetchRelease(slug string) (module model.Module, err error) {
	err = executeRequest(fmt.Sprintf("%s/v3/releases/%s", forge.URI, slug), &module)
	return
}

// ListReleasePlans Returns a paginated list of all plans from the module release identified by the release_slug name.
// The release_slug is composed of the hyphenated module author, name, and version number (example: puppetlabs-lvm-1.4.0).
func (forge *ForgeReleaseBridge) ListReleasePlans(slug string) (module model.Module, err error) {
	err = executeRequest(fmt.Sprintf("%s/v3/releases/%s/plans", forge.URI, slug), &module)
	return
}

// FetchReleasePlan Returns a summary of the given plan from the module release identified by the release_slug name.
// The release_slug is composed of the hyphenated module author, name, and version number (example: puppetlabs-lvm-1.4.0).
// The plan_name should be the full name including the module name (example: lvm::expand).
func (forge *ForgeReleaseBridge) FetchReleasePlan(slug string, plan string) (module model.Module, err error) {
	err = executeRequest(fmt.Sprintf("%s/v3/releases/%s/plans/%s", forge.URI, slug, plan), &module)
	return
}

// DownloadRelease Download a module release tarball
func (forge *ForgeReleaseBridge) DownloadRelease(filename string) (module model.Module, err error) {
	err = executeRequest(fmt.Sprintf("%s/v3/files/%s", forge.URI, filename), &module)
	return
}
