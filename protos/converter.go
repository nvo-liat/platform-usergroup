package protos

import (
	"github.com/nvo-liat/platform-usergroup/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertUsergroupResponse(m *UsergroupResponse) (mx *entity.Usergroup, e error) {
	mx, e = ConvertUsergroupToEntity(m.Usergroup)

	return
}

func ConvertUsergroupToEntity(m *Usergroup) (mx *entity.Usergroup, e error) {
	id, e := primitive.ObjectIDFromHex(m.Id)
	if e != nil {
		return
	}

	mx = &entity.Usergroup{
		ID:         id,
		ClientID:   m.ClientId,
		Name:       m.Name,
		Privileges: m.Privileges,
	}

	return
}

func ConvertUsergroup(m *entity.Usergroup) (mx *Usergroup) {

	return &Usergroup{
		Id:         m.ID.Hex(),
		ClientId:   m.ClientID,
		Name:       m.Name,
		Privileges: m.Privileges,
	}
}
