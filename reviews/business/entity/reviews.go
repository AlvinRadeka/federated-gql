package entity

// Review ...
type Review struct {
	Body       string
	ProductID  string
	ReviewerID string
}

// Product ...
type Product struct {
	ID    string
	Name  string
	Price float64
}

// User ...
type User struct {
	ID   string
	Name string
}
