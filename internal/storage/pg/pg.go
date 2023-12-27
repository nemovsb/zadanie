package pg

import (
	"fmt"
	"log"
	"os"
	"time"
	"zadanie/internal/domain"

	uuid "github.com/satori/go.uuid"
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

	err = db.AutoMigrate(&Goods{}, &Warehouses{})
	if err != nil {
		return nil, fmt.Errorf("migrate gorm models error: %w", err)
	}

	return &Postgres{db: db}, nil
}

func (pg *Postgres) ReserveGoods(goodsIDs []string) error {
	reserveID := uuid.NewV4().String()

	tx := pg.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, goodId := range goodsIDs {
		err := tx.Raw(`
		UPDATE goods
		SET reserve=?
		WHERE id=?
		`,
			reserveID,
			goodId,
		).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (pg *Postgres) ReleaseGoods(goodsIDs []string) error {
	tx := pg.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	for _, goodId := range goodsIDs {
		err := tx.Raw(`
		UPDATE goods
		SET reserve=NULL
		WHERE id=?
		`,
			goodId,
		).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}

func (pg *Postgres) GetRemainGoods(warehouseID string) ([]domain.Good, error) {

	tx := pg.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var res []Goods

	err := tx.Table("goods").Where("wh_id=?", warehouseID).Find()

	return nil, nil
}
