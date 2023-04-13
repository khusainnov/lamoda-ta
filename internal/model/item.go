package model

// Item struct contains info about every item
type Item struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Size        string `json:"size,omitempty"`
	Code        string `json:"code,omitempty"`
	Count       int    `json:"count,omitempty"`
	WarehouseID int    `json:"warehouse_id,omitempty"`
}

// Reservation struct contains which item were reserved
type Reservation struct {
	ItemID string `json:"item_id,omitempty"`
	Count  int    `json:"count,omitempty"`
}
