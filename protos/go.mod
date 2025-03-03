module github.com/nvo-liat/platform-usergroup/protos

go 1.23

toolchain go1.23.0

require (
	github.com/asim/go-micro/v3 v3.7.1
	github.com/nvo-liat/platform-usergroup/entity v0.0.2-develop-622d8e0
	go.mongodb.org/mongo-driver v1.17.3
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/miekg/dns v1.1.63 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
)

replace github.com/nvo-liat/platform-usergroup/entity => ../entity
