# OASIS-data Module

This module contains the data access layer for the OASIS platform. It enforces a clean architecture by isolating domain models, abstracting database interactions through interfaces, and providing concrete database implementations.

## Directory Structure

```text
oasis-data/
├── models/
│   └── hr/
│       ├── department_model.go
│       └── employee_model.go
├── interfaces/
│   ├── department_interface.go
│   └── employee_interface.go
└── adapters/
    ├── postgres/
    │   ├── postgres_department_adapter.go
    │   └── postgres_employee_adapter.go
    └── sqlite/
        ├── sqlite_department_adapter.go
        └── sqlite_employee_adapter.go
```

## Components

### `models/`
Contains the core domain structures divided by business domain (e.g., `hr/`). These structs define the shape of the data exclusively and do not contain any database tags or framework-specific logic, ensuring the domain layer remains agnostic.

### `interfaces/`
Holds the definition contracts (repositories) for how data should be retrieved and manipulated. Modules containing business logic or gRPC endpoints rely purely on these interfaces rather than concrete implementations.
- Utilizes keyset (token-based) pagination for scalable `List` operations.
- Enforces batch query patterns like `BatchGet`.

### `adapters/`
Contains the concrete implementations of the `interfaces` using distinct database technologies. By hiding database logic under adapters, the rest of the ecosystem can mock, swap, or upgrade databases without modifying business code.
- **`postgres/`**: Intended for the production environment utilizing PostgreSQL.
- **`sqlite/`**: Intended for local development, MVP, or testing scenarios utilizing SQLite.

## Usage

When bootstrapping the main application server, you will initialize the adapter configuration of your choice and inject it into the layers that depend on it:

```go
import (
    "oasis-data/interfaces"
    "oasis-data/adapters/postgres"
)

func main() {
    // 1. Initialize concrete database repository (Adapter)
    var employeeRepo interfaces.EmployeeInterface = postgres.NewPostgresEmployeeAdapter()

    // 2. Inject repository into business logic or handlers
    // employee, err := employeeRepo.GetByID("1234")
}
```
