package src

import (
	"github.com/nvo-liat/platform-usergroup/protos"
	"github.com/nvo-liat/platform-usergroup/src/handler/grpc"
	"github.com/nvo-liat/platform-usergroup/src/handler/rest"

	"github.com/asim/go-micro/v3"
	"github.com/labstack/echo/v4"
)

func RegisterRestHandler(e *echo.Echo) {
	rest.RegisterHandler(e)
}

func RegisterGrpcHandler(s micro.Service) {
	protos.RegisterUsergroupServiceHandler(s.Server(), new(grpc.UsergroupService))
}
