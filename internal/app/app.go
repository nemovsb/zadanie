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

func (a *App) ReserveGoods(goodsIDs []string) error {
	if err := a.Storage.ReserveGoods(goodsIDs); err != nil {
		return fmt.Errorf("reserve goods error: %w", err)
	}

	return nil
}

func (a *App) ReleaseGoods(goodsIDs []string) error {
	if err := a.Storage.ReserveGoods(goodsIDs); err != nil {
		return fmt.Errorf("release goods error: %w", err)
	}

	return nil
}

func (a *App) GetRemainGoods(warehouseID string) ([]domain.Good, error) {
	return nil, nil
}
