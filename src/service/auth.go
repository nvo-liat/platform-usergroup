package service

import (
	"context"

	"github.com/env-io/factory/grpc"
	"github.com/nvo-liat/platform-auth/entity"
	"github.com/nvo-liat/platform-auth/protos"
)

type AuthService struct {
	srv protos.AuthService
}

func NewAuthService() *AuthService {
	return &AuthService{
		srv: protos.NewAuthService(grpc.Service.Client()),
	}
}

func (s *AuthService) Session(id, action string) (sd *entity.SessionData, e error) {
	resp, e := s.srv.Session(context.TODO(), &protos.SessionRequest{Id: id, Service: "liat.platform.usergroup", Action: action})
	if e == nil {
		sd = protos.ConvertSessionResponse(resp)
	}

	return
}

func (s *AuthService) ShowAuthorization(id string) (mx *entity.Authorization, e error) {
	resp, e := s.srv.ShowAuthorization(context.TODO(), &protos.ShowRequest{Id: id})
	if e == nil {
		mx, e = protos.ConvertAuthorizationResponse(resp)
	}

	return
}
