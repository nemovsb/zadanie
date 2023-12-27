package pg

import (
	"fmt"
	"log"
	"os"
	"time"
	"zadanie/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("can't create db connection: %w", err)
	}

	err = db.AutoMigrate(&Goods{}, &Warehouses{}, &Reserve{})
	if err != nil {
		return nil, fmt.Errorf("migrate gorm models error: %w", err)
	}

	return &Postgres{db: db}, nil
}

func (pg *Postgres) ReserveGoods(goodsIDs []string) error                     { return nil }
func (pg *Postgres) ReleaseGoods(goodsIDs []string) error                     { return nil }
func (pg *Postgres) GetRemainGoods(warehouseID string) ([]domain.Good, error) { return nil, nil }
