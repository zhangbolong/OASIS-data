package hr

// Department represents an organizational unit within the company.
type Department struct {
	ID          string
	Name        string
	ParentID    string // Optional: for hierarchy
	Description string // Optional description
}
