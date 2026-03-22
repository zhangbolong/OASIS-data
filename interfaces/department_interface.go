package interfaces

import "oasis-data/models/hr"

// DepartmentInterface defines database-agnostic operations for Department models.
// It includes standard CRUD operations that concrete adapters should implement.
type DepartmentInterface interface {
	GetByID(id string) (*hr.Department, error)
	// List retrieves Departments using keyset pagination (token-based).
	// The token encodes the position (e.g., the last seen ID) to start the next page.
	List(limit int, token string) (departments []*hr.Department, nextToken string, err error)
	BatchGet(ids []string) ([]*hr.Department, error)
	Create(department *hr.Department) error
	Update(department *hr.Department) error
	// UpdateDepartmentHierarchy atomicly updates multiple department records for hierarchy changes
	UpdateDepartmentHierarchy(departments []*hr.Department) error
	Delete(id string) error
}
