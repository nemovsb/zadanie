package storage_mock

import (
	"fmt"
	"zadanie/internal/domain"
)

type StorageMock struct{}

func NewStorageMock() *StorageMock {
	return &StorageMock{}
}

func (m *StorageMock) ReserveGoods(goodsIDs []int64) (string, error) {
	if len(goodsIDs) == 0 {
		return "", fmt.Errorf("error")
	}

	if goodsIDs[0] == 0 {
		return "", fmt.Errorf("error")
	}

	if goodsIDs[0] == 1 {
		return "test_reserve_id_1", nil
	}

	return "", fmt.Errorf("error")
}

func (m *StorageMock) ReleaseGoods(goodsIDs []int64) error {
	if len(goodsIDs) <= 0 {
		return fmt.Errorf("empty goods array")
	}

	if goodsIDs[0] <= 0 {
		return fmt.Errorf("wrong good id")
	}

	if goodsIDs[0] == 1 {
		return nil
	}

	return fmt.Errorf("wrong good ID")
}

func (m *StorageMock) GetRemainGoods(warehouseID int64) ([]domain.Good, error) {

	if warehouseID <= 0 {
		return nil, fmt.Errorf("wrong wh ID")
	}

	if warehouseID == 1 {
		return []domain.Good{{
			ID:       1,
			Name:     "name1",
			Size:     "1x1x1",
			Quantity: 10},
		}, nil
	}

	return nil, fmt.Errorf("no warehouse")
}
