package main

import (
	"os"
	"time"

	"github.com/nvo-liat/platform-usergroup/protos"
	"github.com/nvo-liat/platform-usergroup/src"

	"github.com/env-io/factory"
	"github.com/env-io/factory/grpc"
	"github.com/env-io/factory/kafka"
	"github.com/env-io/factory/mongo"
	"github.com/env-io/factory/rest"
	"github.com/joho/godotenv"
)

func init() {
	// load env variable file if exists
	godotenv.Load()

	// application config
	factory.AppConfig = &factory.Config{
		AppName:    protos.ServiceName,
		IsDev:      os.Getenv("DEBUG_MODE") == "true",
		AppVersion: os.Getenv("APP_VERSION"),
		AppService: os.Getenv("APP_SERVICE"),
	}

	// new logger instances
	factory.NewLogger(factory.AppConfig.AppName)

	// new mongo instances
	initMongoConnection()

	// new kafka connection
	initKafkaConnection()

	// initial rest server
	initRestServer()

	// initial grpc server
	initGrpcServer()
}

// @title liat.platform.usergroup
// @description host https://api.liat.co.id/sandbox/usergroup
// @version v1
// @host https://api.liat.co.id/sandbox/usergroup
// @BasePath /
func main() {
	factory.Routine.Add(rest.Start, rest.Shutdown)
	factory.Routine.Add(grpc.Start, grpc.Shutdown)

	factory.Logger.Sugar().Error(factory.Routine.Run())
}

func initMongoConnection() {
	c := &mongo.Config{
		Server:     os.Getenv("MONGO_SERVER"),
		Username:   os.Getenv("MONGO_USERNAME"),
		Password:   os.Getenv("MONGO_PASSWORD"),
		Database:   os.Getenv("MONGO_DATABASE"),
		Datasource: os.Getenv("MONGO_DATASOURCE"),
		CtxTimeout: 1 * time.Minute,
	}

	if e := mongo.NewConnection(c); e != nil {
		factory.Logger.Error(e.Error())
	}
}

func initRestServer() {
	c := &rest.Config{
		Server:    os.Getenv("REST_SERVER"),
		IsDev:     factory.AppConfig.IsDev,
		JwtSecret: os.Getenv("REST_JWT"),
	}

	rest.NewServer(c, src.RegisterRestHandler)
}

func initGrpcServer() {
	c := &grpc.Config{
		Name:           factory.AppConfig.AppName,
		RegistryServer: os.Getenv("SERVICE_REGISTRY"),
		Server:         os.Getenv("SERVICE_SERVER"),
	}

	if e := grpc.NewService(c, src.RegisterGrpcHandler); e != nil {
		factory.Logger.Panic(e.Error())
	}
}

func initKafkaConnection() {
	c := &kafka.Config{
		Server: os.Getenv("KAFKA_SERVER"),
		Name:   factory.AppConfig.AppName,
	}

	if e := kafka.NewConnection(c, nil); e != nil {
		factory.Logger.Panic(e.Error())
	}
}
