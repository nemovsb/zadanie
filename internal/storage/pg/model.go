package pg

import "github.com/jackc/pgtype"

type Goods struct {
	ID      int64        `gorm:"column:id;type:bigint;primaryKey:PK1"`
	Name    string       `gorm:"column:name;type:varchar(100);not null"`
	Size    string       `gorm:"column:size;type:varchar(100);default:'1x1x1';not null"`
	Qantity int64        `gorm:"column:quantity;type:bigint"`
	Reserve pgtype.JSONB `gorm:"column:reserve;type:jsonb;default:'[]';not null"`
}

type Warehouses struct {
	ID     int64 `gorm:"column:id;type:bigint;primaryKey"`
	Status bool  `gorm:"column:status"`
}

type Reserve struct {
	ID          string     `gorm:"column:id;type:varchar(50);primaryKey:PK1"`
	Goods       Goods      `gorm:"foreignKey:GoodsID;references:ID"`
	GoodsID     int64      `gorm:"column:goods_id;type:bigint;not null"`
	Warehouses  Warehouses `gorm:"foreignKey:WarehouseID;refernces:ID"`
	WarehouseID int64      `gorm:"column:wh_id;type:bigint;not null"`
	Quantity    int64      `gorm:"column:quantity;type:bigint;not null"`
}
