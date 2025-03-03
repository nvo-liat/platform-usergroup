module github.com/nvo-liat/platform-usergroup/protos

go 1.23

toolchain go1.23.0

require (
	github.com/asim/go-micro/v3 v3.7.1
	github.com/nvo-liat/platform-usergroup/entity v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.17.1
	google.golang.org/protobuf v1.35.2
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.0.0-20210502180810-71e4cd670f79 // indirect
	golang.org/x/text v0.17.0 // indirect
)

replace github.com/nvo-liat/platform-usergroup/entity => ../entity
