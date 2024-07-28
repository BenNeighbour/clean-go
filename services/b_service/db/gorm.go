package db

import (
	"benneighbour.com/services/b_service/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Register initializes a new Gorm DB instance with a MySQL connection.
func Register() *gorm.DB {
	// Construct the MySQL connection string.
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		*config.DB_USERNAME,
		*config.DB_PASSWORD,
		*config.DB_URL,
		*config.DB_PORT,
		*config.DB_NAME,
	)

	// Open the database connection.
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Retrieve the underlying sql.DB object to configure connection pool.
	connection, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting DB object: %v", err)
	}

	// Configure the connection pool.
	connection.SetMaxIdleConns(*config.DB_POOL_IDLE)
	connection.SetMaxOpenConns(*config.DB_POOL_MAX)
	connection.SetConnMaxLifetime(*config.DB_POOL_LIFETIME)

	return db
}
