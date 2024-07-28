package config

import (
	"flag"
	"time"
)

// Define flags for server configuration.
var (
	// Server Variables
	SERVER_PORT = flag.Int("port", 7777, "The server port")

	// DB Connection Variables
	DB_PORT     = flag.Int("db_port", 3306, "The MySQL DB port")
	DB_URL      = flag.String("db_url", "localhost", "The MySQL DB URL")
	DB_USERNAME = flag.String("db_username", "root", "The MySQL DB username")
	DB_PASSWORD = flag.String("db_password", "root", "The MySQL DB password")
	DB_NAME     = flag.String("db_name", "b_service", "The MySQL DB name")

	// DB Pool Variables
	DB_POOL_LIFETIME = flag.Duration("db_pool_lifetime", 3000*time.Millisecond, "The maximum lifetime of a connection to the MySQL database in the pool")
	DB_POOL_MAX      = flag.Int("db_pool_max", 100, "The maximum number of open connections to the MySQL database")
	DB_POOL_IDLE     = flag.Int("db_pool_idle", 5, "The maximum number of idle connections in the MySQL database pool")

	// TODO - Logging variables
)

func Initialize() {
	// Parse command-line flags to initialize the environment.
	flag.Parse()
}
