package app

import (
	"context"
	"fmt"
	"zadanie/internal/domain"
)

type App struct {
	Storage
}

func NewApp(s Storage) *App {
	return &App{
		Storage: s,
	}
}

func (a *App) Run(ctx context.Context) error {
	select {
	case <-ctx.Done():
		{
			return nil
		}
	}
}

func (a *App) Shutdown(cancel context.CancelFunc) error {
	cancel()
	return nil
}

func (a *App) ReserveGoods(goodsIDs []int64) (string, error) {
	reserveID, err := a.Storage.ReserveGoods(goodsIDs)
	if err != nil {
		return "", fmt.Errorf("reserve goods error: %w", err)
	}

	return reserveID, nil
}

func (a *App) ReleaseGoods(goodsIDs []int64) error {
	if err := a.Storage.ReleaseGoods(goodsIDs); err != nil {
		return fmt.Errorf("release goods error: %w", err)
	}

	return nil
}

func (a *App) GetRemainGoods(warehouseID int64) ([]domain.Good, error) {
	goods, err := a.Storage.GetRemainGoods(warehouseID)
	if err != nil {
		return nil, fmt.Errorf("get free goods error: %w", err)
	}
	return goods, nil
}
