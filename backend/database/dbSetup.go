package database

import (
	"github.com/simpleittools/simplehome/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	//Postgres
	dsn := "host=localhost user=home_admin password=mONK!ion48 dbname=simplehomedb port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// SQLite (would need to replace postgres driver with "gorm.io/driver/sqlite")
	//conn, err := gorm.Open(sqlite.Open("simplehome.db"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the DB")
	}

	DB = conn

	// AutoMigrate will migrate the models in to the table
	conn.AutoMigrate(&models.User{})

}
