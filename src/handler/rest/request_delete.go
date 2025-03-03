package rest

import (
	"github.com/nvo-liat/platform-usergroup/entity"
	"github.com/nvo-liat/platform-usergroup/src/bloc"

	"github.com/env-io/validate"
	auth "github.com/nvo-liat/platform-auth/entity"
)

type deleteRequest struct {
	ID string `json:"-"`

	Session   *auth.SessionData `json:"-"`
	Usergroup *entity.Usergroup `json:"-"`
}

func (r *deleteRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	var e error

	if r.ID != "" {
		if r.Usergroup, e = bloc.ValidID(r.ID); e != nil {
			v.SetError("id.invalid", "data tidak ditemukan.")
		}
	}

	if r.Usergroup != nil {
		if r.Usergroup.IsDeleted {
			v.SetError("id.invalid", "data sudah dihapus.")
		}
	}

	return v
}

func (r *deleteRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *deleteRequest) Execute() (e error) {
	return bloc.UsergroupDeleting(r.Usergroup, r.Session)

}
