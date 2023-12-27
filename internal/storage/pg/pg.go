package pg

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(config *Config) (*Postgres, error) {
	dsn := fmt.Sprintf(`
	host=%s
	port=%s
	user=%s
	password=%s
	dbname=%s
	sslmode=disable
	TimeZone=UTC`,
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("can't create db connection: %w", err)
	}

	return &Postgres{db: db}, nil
}
