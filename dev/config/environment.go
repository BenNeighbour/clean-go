package config

import (
	"flag"
)

// Define flags for server configuration.
var (
	SERVER_PORT = flag.Int("dev_port", 5555, "The server port")
)

func Initialize() {
	// Parse command-line flags to initialize the environment.
	flag.Parse()
}
