package hr

// Employee represents an individual working within the organization.
// It includes core HR details for future domains and reporting structures.
type Employee struct {
	ID            string // Unique employee ID
	FirstName     string
	LastName      string
	Email         string
	PhoneNumber   string
	DepartmentID  string   // References Department.ID
	Role          string   // Job title or position
	Level         int      // Rank or seniority level
	Status        string   // Active, Terminated, OnLeave, etc.
	StartDate     string   // Could be time.Time in future
	EndDate       string   // Optional, for terminated employees
	ManagerID     string   // Employee.ID of direct reporting manager, empty if top-level
	DirectReports []string // Optional: list of Employee.IDs who report directly
}
