package model

type ResponseMsg string

const (
	RequestErr ResponseMsg = "error due processing request"
	RequestOk  ResponseMsg = "success"
)

// Product struct contains info about every item
type Product struct {
	ID          string `json:"id,omitempty" db:"id"`
	Name        string `json:"name,omitempty" db:"name"`
	Size        string `json:"size,omitempty" db:"size"`
	Code        string `json:"code,omitempty" db:"code"`
	Quantity    int    `json:"quantity,omitempty" db:"quantity"`
	WarehouseID string `json:"warehouse_id,omitempty" db:"warehouse_id"`
}

// Reservation struct contains which item were reserved
type Reservation struct {
	ID    string   `json:"id,omitempty"`
	Code  []string `json:"code,omitempty"`
	Count int      `json:"count,omitempty"`
}

type UploadProductRequest struct {
	Name        string `json:"name,omitempty"`
	Size        string `json:"size,omitempty"`
	Code        string `json:"code,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	WarehouseID string `json:"warehouse_id,omitempty"`
}

type UpdateProductRequest struct {
	Name        string `json:"name,omitempty"`
	Size        string `json:"size,omitempty"`
	Code        string `json:"code,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	WarehouseID string `json:"warehouse_id,omitempty"`
}

type GetProductRequest struct {
	ProductCode string `json:"product_code"`
	WarehouseID string `json:"warehouse_id"`
}

type ListProductRequest struct {
	WarehouseID string `json:"warehouse_id"`
}

type DeleteProductRequest struct {
	ProductCode string `json:"product_code"`
	WarehouseID string `json:"warehouse_id"`
}

type ReserveProductRequest struct {
	ProductCode []string `json:"product_code"`
	WarehouseID string   `json:"warehouse_id"`
	Quantity    int      `json:"quantity"`
}

type ReleaseReserveRequest struct {
	ProductCode []string `json:"product_code"`
	WarehouseID string   `json:"warehouse_id"`
}

type UploadProductResponse struct {
	Message ResponseMsg `json:"message,omitempty"`
}

type Response struct {
	Message ResponseMsg `json:"message,omitempty"`
}
