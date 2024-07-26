package config

import (
	"google.golang.org/grpc"

	a "benneighbour.com/services/a_service/transport"
	b "benneighbour.com/services/b_service/transport"
	//c "benneighbour.com/services/c_service/transport"
	//d "benneighbour.com/services/d_service/transport"
	//e "benneighbour.com/services/e_service/transport"
	//f "benneighbour.com/services/f_service/transport"
)

var Services = []func(*grpc.Server){
	a.Register,
	b.Register,
	//{Name: "c_transport", Register: c.Register},
	//{Name: "d_transport", Register: d.Register},
	//{Name: "e_transport", Register: e.Register},
	//{Name: "f_transport", Register: f.Register},
	//etc...
}
