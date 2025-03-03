package repository

import (
	"github.com/nvo-liat/platform-usergroup/entity"

	"github.com/env-io/factory/helper"
	"github.com/env-io/factory/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsergroupRepository struct {
	coll *mongo.Collection
}

func NewUsergroupRepository() *UsergroupRepository {
	return &UsergroupRepository{
		coll: mongo.NewCollection("usergroup"),
	}
}

func (r *UsergroupRepository) Get(query []primitive.M, req helper.GetRequest) (result entity.Usergroups, total int64, e error) {
	filter := bson.M{}

	if req != nil {
		if s := req.GetSearch(); s != "" {
			query = append(query, mongo.FilterSearch(s, "name"))
		}
	}

	if len(query) > 0 {
		filter = bson.M{"$and": query}
	}

	if total, e = r.coll.Counts(filter); e != nil {
		return
	}

	options := options.Find()
	if req != nil {
		options.SetSort(mongo.RequestSort(req.GetOrders()))
		options.SetLimit(req.GetLimit())
		options.SetSkip(req.GetOffset())
	}

	e = r.coll.Finds(&result, filter, options)

	return
}

func (r *UsergroupRepository) Create(req *entity.Usergroup) (e error) {
	return r.coll.Create(req)
}

func (r *UsergroupRepository) Show(id string) (mx *entity.Usergroup, e error) {
	e = r.coll.Show(id, &mx)

	return
}

func (r *UsergroupRepository) Update(req *entity.Usergroup, fields ...string) (e error) {
	return r.coll.Update(req, fields...)
}

func (r *UsergroupRepository) Delete(req *entity.Usergroup) (e error) {
	req.IsDeleted = true

	return r.coll.Update(req)
}

func (r *UsergroupRepository) FindByName(name string, exid string) (m *entity.Usergroup, e error) {
	query := []bson.M{{"name": name}, {"is_deleted": false}}

	if exid != "" {
		var exculde primitive.ObjectID
		if exculde, e = primitive.ObjectIDFromHex(exid); e != nil {
			return
		}

		query = append(query, bson.M{"_id": bson.M{"$ne": exculde}})
	}

	e = r.coll.GetOne(nil, bson.M{"$and": query}, &m)

	return
}
