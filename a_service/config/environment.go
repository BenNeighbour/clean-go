package config

import (
	"flag"
)

var (
	SERVER_PORT = flag.Int("port", 6666, "The server port")
	DB_PORT     = flag.Int("db_port", 5432, "The cassandra db port")
	DB_URL      = flag.Int("db_url", 5555, "The cassandra db url")
	DB_USERNAME = flag.Int("db_username", 5555, "The cassandra db username")
	DB_PASSWORD = flag.Int("db_password", 5555, "The cassandra db password")
)
