package bloc

import (
	"github.com/nvo-liat/platform-usergroup/entity"
	"github.com/nvo-liat/platform-usergroup/src/repository"

	auth "github.com/nvo-liat/platform-auth/entity"
)

func UsergroupCreating(m *entity.Usergroup, su *auth.SessionData) (mx *entity.Usergroup, e error) {

	if e = repository.NewUsergroupRepository().Create(m); e != nil {
		return
	}

	// go publisher.NewUsergroupEvent(m, su).Created()

	return m, e
}

func UsergroupUpdating(m *entity.Usergroup, fields []string, su *auth.SessionData) (mx *entity.Usergroup, e error) {
	if e = repository.NewUsergroupRepository().Update(m, fields...); e != nil {
		return
	}

	// go publisher.NewUsergroupEvent(m, su).Updated()

	return m, e
}

func UsergroupDeleting(m *entity.Usergroup, su *auth.SessionData) (e error) {
	if e = repository.NewUsergroupRepository().Delete(m); e != nil {
		return
	}

	// go publisher.NewUsergroupEvent(m, su).Deleted()

	return
}
