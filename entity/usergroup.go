package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usergroup struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ClientID   string             `bson:"client_id" json:"client_id"`
	Name       string             `bson:"name" json:"name"`
	Privileges []string           `bson:"privileges" json:"privileges"`
	IsDeleted  bool               `bson:"is_deleted" json:"-"`
}

type Usergroups []*Usergroup
