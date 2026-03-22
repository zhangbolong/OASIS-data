package hr

// Department represents an organizational unit within the company.
type Department struct {
	ID          string
	Name        string
	ParentID    string   // Optional: for hierarchy
	ChildIDs    []string // Downward links for hierarchy
	Description string   // Optional description
}

// DepartmentBrief represents a summarized view of a Department.
type DepartmentBrief struct {
	ID   string
	Name string
}
