package domain

type Good struct {
	ID       string
	Name     string
	Size     string
	Quantity uint32
}

type Warehouse struct {
	Name         string
	Availability bool
}
