module benneighbour.com/dev

go 1.22.5

require (
	benneighbour.com/services/a_service v0.0.0-00010101000000-000000000000
	benneighbour.com/services/b_service v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
)

require (
	benneighbour.com/proto v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240723171418-e6d459c13d2a // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace benneighbour.com/services/b_service => ../services/b_service

replace benneighbour.com/services/a_service => ../services/a_service

replace benneighbour.com/proto => ../proto
