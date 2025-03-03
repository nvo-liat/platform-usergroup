package rest

import (
	"github.com/nvo-liat/platform-usergroup/src/repository"

	"github.com/env-io/factory/helper"
	"github.com/env-io/factory/rest"

	auth "github.com/nvo-liat/platform-auth/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type getRequest struct {
	helper.DefaultGetRequest

	Session *auth.SessionData `json:"-"`
}

func (r *getRequest) Detail(id string) (resp *rest.ResponseBody) {
	resp = rest.NewResponseBody()

	if u, e := repository.NewUsergroupRepository().Show(id); e == nil {
		resp.Body(u, 0)
	}

	return
}

func (r *getRequest) List() (resp *rest.ResponseBody) {
	resp = rest.NewResponseBody()

	query := []bson.M{{"is_deleted": false}}

	if mx, total, e := repository.NewUsergroupRepository().Get(query, r); e == nil && total != 0 {
		resp.Body(mx, total)
	}

	return
}
