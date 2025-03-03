package bloc

import (
	"github.com/nvo-liat/platform-usergroup/entity"
	"github.com/nvo-liat/platform-usergroup/src/repository"
	"github.com/nvo-liat/platform-usergroup/src/service"
)

func ValidUniqueUsergroup(name string, exclude string) bool {
	_, err := repository.NewUsergroupRepository().FindByName(name, exclude)

	return err != nil
}

func ValidAuthorizationID(id string) (e error) {
	_, e = service.NewAuthService().ShowAuthorization(id)

	return
}

func ValidID(id string) (mx *entity.Usergroup, e error) {
	mx, e = repository.NewUsergroupRepository().Show(id)

	return
}
