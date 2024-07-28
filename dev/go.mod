module benneighbour.com/dev

go 1.22.5

require (
	benneighbour.com/services/a_service v0.0.0-00010101000000-000000000000
	benneighbour.com/services/b_service v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
)

require (
	benneighbour.com/proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240723171418-e6d459c13d2a // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	gorm.io/gorm v1.25.11 // indirect
)

replace benneighbour.com/services/b_service => ../services/b_service

replace benneighbour.com/services/a_service => ../services/a_service

replace benneighbour.com/proto => ../proto
