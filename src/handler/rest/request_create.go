package rest

import (
	"fmt"

	"github.com/nvo-liat/platform-usergroup/entity"
	"github.com/nvo-liat/platform-usergroup/src/bloc"

	"github.com/env-io/validate"
	auth "github.com/nvo-liat/platform-auth/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createRequest struct {
	Name       string   `json:"name" valid:"required"`
	Privileges []string `json:"privileges" valid:"required"`

	Session *auth.SessionData `json:"-"`
}

func (r *createRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	if r.Name != "" {
		if !bloc.ValidUniqueUsergroup(r.Name, "") {
			v.SetError("name.invalid", "nama sudah tersedia.")
		}
	}

	if len(r.Privileges) > 0 {
		for key, i := range r.Privileges {
			var e error
			if _, e = primitive.ObjectIDFromHex(i); e != nil {
				v.SetError(fmt.Sprintf("privileges.%d.privilege.invalid", key), "privilege tidak valid.")
			}

			if e = bloc.ValidAuthorizationID(i); e != nil {
				v.SetError(fmt.Sprintf("privileges.%d.privilege.invalid", key), "privilege tidak ditemukan.")
			}
		}
	}

	return v
}

func (r *createRequest) Messages() map[string]string {
	return map[string]string{
		"name.required":       "nama harus diisi.",
		"privileges.required": "aksi usergroup harus diisi.",
	}
}

func (r *createRequest) Execute() (m *entity.Usergroup, e error) {
	m = &entity.Usergroup{
		Name:      r.Name,
		IsDeleted: false,
	}

	if len(r.Privileges) > 0 {
		for _, id := range r.Privileges {
			m.Privileges = append(m.Privileges, id)
		}
	}

	m, e = bloc.UsergroupCreating(m, r.Session)

	return
}
