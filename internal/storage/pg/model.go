package pg

type Goods struct {
	ID          uint       `gorm:"column:id;primaryKey:PK1;autoincrement"`
	Name        string     `gorm:"column:name;type:varchar(100);not null"`
	Size        string     `gorm:"column:size;type:varchar(100);default:'1x1x1';not null"`
	Qantity     int64      `gorm:"column:quantity;type:bigint"`
	Warehouses  Warehouses `gorm:"foreignKey:WarehouseID;refernces:ID"`
	WarehouseID uint       `gorm:"column:wh_id;not null"`
	ReserveID   string     `gorm:"column:reserve_id;type:varchar(50)"`
}

type Warehouses struct {
	ID     uint `gorm:"column:id;primaryKey;autoincrement"`
	Status bool `gorm:"column:status"`
}

// type Stock struct {
// 	ID          int64      `gorm:"column:id;type:serial;primaryKey:PK1"`
// 	Goods       Goods      `gorm:"foreignKey:GoodsID;references:ID"`
// 	GoodsID     int64      `gorm:"column:goods_id;type:serial;not null"`
// 	Warehouses  Warehouses `gorm:"foreignKey:WarehouseID;refernces:ID"`
// 	WarehouseID int64      `gorm:"column:wh_id;type:serial;not null"`
// 	Quantity    int64      `gorm:"column:quantity;type:bigint;not null"`
// }

// type Reserve struct {
// 	ID       string `gorm:"column:id;type:varchar(50);primaryKey:PK1"`
// 	Stock    Stock  `gorm:"foreignKey:StockID;references:ID"`
// 	StockID  int64  `gorm:"column:stock_id;type:serial;not null"`
// 	Quantity int64  `gorm:"column:quantity;type:bigint;not null"`
// }
