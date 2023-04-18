package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/lamoda/internal/model"
)

// Product interface contains methods for working with items
type Product interface {
	UploadProduct(product model.UploadProductRequest) error
	UpdateProduct(data model.UpdateProductRequest) error
	GetProduct(data model.GetProductRequest) (*model.Product, error)
	ListProduct(data model.ListProductRequest) (*[]model.Product, error)
	DeleteProduct(data model.DeleteProductRequest) error
}

// Reservation interface contains methods for working with item reservation
type Reservation interface {
	ReserveProduct(data model.ReserveProductRequest) error
	ReleaseReserve(data model.ReleaseReserveRequest) error
}

type Repository struct {
	Product
	Reservation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Product:     NewProductRepository(db),
		Reservation: NewReservationRepository(db),
	}
}
