package domain

type Car struct {
	ID      uint   `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint   `json:"price"`
	Status  Status `json:"status"`
	Mileage uint   `json:"mileage"`
}

// Status shows the cars sale status
type Status string

const (
	InTransit    Status = "in transit"
	InStock      Status = "in stock"
	Sold         Status = "sold"
	Discontinued Status = "discontinued"
)
