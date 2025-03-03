package rest

import (
	"fmt"

	"github.com/nvo-liat/platform-usergroup/entity"
	"github.com/nvo-liat/platform-usergroup/src/bloc"

	"github.com/env-io/validate"
	auth "github.com/nvo-liat/platform-auth/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type updateRequest struct {
	ID         string   `json:"-"`
	Name       string   `json:"name" valid:"required"`
	Privileges []string `json:"privileges" valid:"required"`

	Session   *auth.SessionData `json:"-"`
	Usergroup *entity.Usergroup `json:"-"`
}

func (r *updateRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	var e error

	if r.ID != "" {
		if r.Usergroup, e = bloc.ValidID(r.ID); e != nil {
			v.SetError("id.invalid", "usergroup tidak ditemukan.")
		}
	}

	if r.Name != "" {
		if !bloc.ValidUniqueUsergroup(r.Name, r.ID) {
			v.SetError("name.invalid", "nama sudah tersedia.")
		}
	}

	if len(r.Privileges) > 0 {
		for key, i := range r.Privileges {
			if _, e := primitive.ObjectIDFromHex(i); e != nil {
				v.SetError(fmt.Sprintf("privileges.%d.privilege.invalid", key), "privilege tidak valid.")
			}

			if e = bloc.ValidAuthorizationID(i); e != nil {
				v.SetError(fmt.Sprintf("privileges.%d.privilege.invalid", key), "privilege tidak ditemukan.")
			}
		}
	}

	return v
}

func (r *updateRequest) Messages() map[string]string {
	return map[string]string{
		"name.required":       "nama harus diisi.",
		"privileges.required": "aksi usergroup harus diisi.",
	}
}

func (r *updateRequest) Execute() (m *entity.Usergroup, e error) {
	m = &entity.Usergroup{
		ID:   r.Usergroup.ID,
		Name: r.Name,
	}

	if len(r.Privileges) > 0 {
		for _, id := range r.Privileges {
			m.Privileges = append(m.Privileges, id)
		}
	}

	fields := []string{"name", "privileges"}

	m, e = bloc.UsergroupUpdating(m, fields, r.Session)

	return
}
