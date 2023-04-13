package model

// Warehouse struct store information about every warehouse
type Warehouse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Available bool   `json:"availability,omitempty"`
}
