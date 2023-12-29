package app

import (
	"context"
	"fmt"
	"zadanie/internal/domain"

	"go.uber.org/zap"
)

type App struct {
	Storage
	logger *zap.Logger
}

func NewApp(s Storage, l *zap.Logger) *App {
	return &App{
		Storage: s,
		logger:  l.With(zap.Namespace("app")),
	}
}

func (a *App) Run(ctx context.Context) error {

	a.logger.Info("application run")

	select {
	case <-ctx.Done():
		{
			a.logger.Info("application done")
			return nil
		}
	}
}

func (a *App) Shutdown(cancel context.CancelFunc) error {
	cancel()

	a.logger.Info("application shutdown")
	return nil
}

func (a *App) ReserveGoods(goodsIDs []int64) (string, error) {

	l := a.logger.With(zap.String("handler", "reserve goods"))
	l.Debug("try reserve goods", zap.Any("goodsIDs", &goodsIDs))

	reserveID, err := a.Storage.ReserveGoods(goodsIDs)
	if err != nil {
		err = fmt.Errorf("reserve goods error: %w", err)
		l.Error("reserve error", zap.Error(err))
		return "", err
	}

	l.Debug("reserve goods success", zap.String("reserve ID", reserveID), zap.Any("goodsIDs", &goodsIDs))
	return reserveID, nil
}

func (a *App) ReleaseGoods(goodsIDs []int64) error {

	l := a.logger.With(zap.String("handler", "release goods"))
	l.Debug("try release goods", zap.Any("goodsIDs", &goodsIDs))

	if err := a.Storage.ReleaseGoods(goodsIDs); err != nil {
		err = fmt.Errorf("release goods error: %w", err)
		l.Error("release error", zap.Error(err))
		return err
	}

	l.Debug("release goods success", zap.Any("goodsIDs", &goodsIDs))
	return nil
}

func (a *App) GetRemainGoods(warehouseID int64) ([]domain.Good, error) {

	l := a.logger.With(zap.String("handler", "rget remain goods"), zap.Int64("warehouseID", warehouseID))
	l.Debug("try get goods")

	goods, err := a.Storage.GetRemainGoods(warehouseID)
	if err != nil {
		err = fmt.Errorf("get free goods error: %w", err)
		l.Error("get free goods for warehouse err", zap.Error(err))
		return nil, err
	}

	l.Debug("get goods", zap.Any("goods", &goods))
	return goods, nil
}
