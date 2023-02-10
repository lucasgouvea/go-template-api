package database

import (
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Start(dbHost string, dbUser string, dbPassword string, dbName string) {
	var err error

	var sb strings.Builder

	sb.WriteString("host=")
	sb.WriteString(dbHost)
	sb.WriteString(" user=")
	sb.WriteString(dbUser)
	sb.WriteString(" password=")
	sb.WriteString(dbPassword)
	sb.WriteString(" dbname=")
	sb.WriteString(dbName)
	sb.WriteString(" port=5432 sslmode=disable")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  sb.String(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func Migrate(models []any) error {
	db := GetDB()
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}
