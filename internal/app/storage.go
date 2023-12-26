package app

import "zadanie/internal/domain"

type Storage interface {
	ReserveGoods(goodsIDs []string) error
	ReleaseGoods(goodsIDs []string) error
	GetRemainGoods(warehouseID string) ([]domain.Good, error)
}
