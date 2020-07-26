package forge

import "github.com/syberalexis/puppetforge-server/pkg/model"

type Database struct {
}

func (databse Database) AddRelease(release model.Release) bool {
	return false
}

func (databse Database) ListModules() ([]model.Module, error) {
	return []model.Module{model.Module{}}, nil
}

func (databse Database) ListReleases() []model.Release {
	return []model.Release{}
}

func (databse Database) GetModule(slug string) model.Module {
	return model.Module{}
}

func (databse Database) GetRelease(slug string) model.Release {
	return model.Release{}
}

func (databse Database) DeleteModule(slug string) bool {
	return false
}

func (databse Database) DeleteRelease(slug string) bool {
	return false
}
