package grpc

import (
	"context"

	"github.com/nvo-liat/platform-usergroup/protos"
	"github.com/nvo-liat/platform-usergroup/src/repository"
)

type UsergroupService struct{}

func (g *UsergroupService) Show(ctx context.Context, req *protos.ShowRequest, resp *protos.UsergroupResponse) error {
	u, e := repository.NewUsergroupRepository().Show(req.Id)

	if e == nil {
		resp.Usergroup = protos.ConvertUsergroup(u)
	}

	return e
}
