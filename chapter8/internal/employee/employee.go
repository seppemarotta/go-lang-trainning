package employee

import (
	"time"
)

type Position int

const (
	Undetermined Position = iota
	Junior
	Senior
	Manager
	CEO
)

type Employee struct {
	ID          int       `json:"id,omitempty"`
	FullName    string    `json:"full_name"`
	Position    Position  `json:"position"`
	Salary      float64   `json:"salary"`
	Joined      time.Time `json:"joined"`
	OnProbation bool      `json:"on_probation"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
