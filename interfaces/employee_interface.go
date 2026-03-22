package interfaces

import "oasis-data/models/hr"

// EmployeeInterface defines database-agnostic operations for Employee models.
// It includes standard CRUD operations that concrete adapters should implement.
type EmployeeInterface interface {
	GetByID(id string) (*hr.Employee, error)
	// List retrieves Employees using keyset pagination (token-based).
	// The token encodes the position (e.g., the last seen ID) to start the next page.
	List(limit int, token string) (employees []*hr.Employee, nextToken string, err error)
	BatchGet(ids []string) ([]*hr.Employee, error)
	Create(employee *hr.Employee) error
	Update(employee *hr.Employee) error
	// UpdateEmployeeHierarchy atomicly updates multiple employee records for hierarchy changes
	UpdateEmployeeHierarchy(employees []*hr.Employee) error
	Delete(id string) error
}
