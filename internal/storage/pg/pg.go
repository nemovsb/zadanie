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

func (pg *Postgres) ReserveGoods(goodsIDs []int64) (string, error) {
	reserveID := uuid.NewV4().String()

	tx := pg.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return "", err
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
			return "", err
		}
	}

	return reserveID, tx.Commit().Error
}

func (pg *Postgres) ReleaseGoods(goodsIDs []int64) error {
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

	return tx.Commit().Error
}

func (pg *Postgres) GetRemainGoods(warehouseID int64) ([]domain.Good, error) {

	tx := pg.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var res []Goods

	if err := tx.Table("goods").Where("wh_id=? AND reserve_id IS NULL", warehouseID).Find(&res).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	freeGoods := make([]domain.Good, 0, len(res))

	for _, val := range res {
		freeGoods = append(freeGoods, domain.Good{
			ID:       val.ID,
			Name:     val.Name,
			Size:     val.Size,
			Quantity: uint32(val.Qantity),
		})
	}

	return freeGoods, nil
}
