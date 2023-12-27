package app

import "zadanie/internal/domain"

type Storage interface {
	ReserveGoods(goodsIDs []int64) (string, error)
	ReleaseGoods(goodsIDs []int64) error
	GetRemainGoods(warehouseID int64) ([]domain.Good, error)
}
