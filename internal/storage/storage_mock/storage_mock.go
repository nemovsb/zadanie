package storage_mock

import "zadanie/internal/domain"

type StorageMock struct{}

func NewStorageMock() *StorageMock {
	return &StorageMock{}
}

func (m *StorageMock) ReserveGoods(goodsIDs []string) error {
	return nil
}

func (m *StorageMock) ReleaseGoods(goodsIDs []string) error {
	return nil
}

func (m *StorageMock) GetRemainGoods(warehouseID string) ([]domain.Good, error) {
	return nil, nil
}
