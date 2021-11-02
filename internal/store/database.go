package store

import (
	"fmt"
	"log"

	"github.com/orlovssky/gread/internal/secrets"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// NewDatabase - Create a new database conection. Schema will be migrated if
// not found
func NewDatabase(s secrets.Secrets) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		s.DBHost, s.DBUser, s.DBPass, s.DBDatabase, s.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	// db.SetMaxOpenConns(100)
	runMigrations(db)
	return db
}

// runMigrations - Runs dataase migrations
func runMigrations(db *gorm.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "schema",
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	n, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}