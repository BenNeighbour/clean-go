package config

import (
	"flag"
)

var (
	SERVER_PORT = flag.Int("port", 5555, "The server port")
)
