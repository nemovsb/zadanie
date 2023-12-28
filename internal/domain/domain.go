package domain

type Good struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Size     string `json:"size"`
	Quantity uint32 `json:"quantity"`
}

type Warehouse struct {
	Name         int64 `json:"name"`
	Availability bool  `json:"availability"`
}
