package domain

type Good struct {
	ID       int64
	Name     string
	Size     string
	Quantity uint32
}

type Warehouse struct {
	Name         int64
	Availability bool
}
