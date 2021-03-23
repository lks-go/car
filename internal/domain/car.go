package domain

type Car struct {
	ID      uint
	Brand   string
	Model   string
	Price   uint
	Status  Status
	Mileage uint
}

// Status shows the cars sale status
type Status string

const (
	InTransit    Status = "in transit"
	InStock      Status = "in stock"
	Sold         Status = "sold"
	Discontinued Status = "discontinued"
)
