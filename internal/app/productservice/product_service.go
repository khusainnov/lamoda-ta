package productservice

import (
	"sync"

	"github.com/khusainnov/lamoda/internal/model"
	"github.com/khusainnov/lamoda/internal/repository"
	"go.uber.org/zap"
)

type Product interface {
	UploadProduct(req model.UploadProductRequest, resp *model.UploadProductResponse) error
	UpdateProduct(req model.UpdateProductRequest, resp *model.Response) error
	GetProduct(req model.GetProductRequest, resp *model.Product) error
	ListProduct(req model.ListProductRequest, resp *[]model.Product) error
	DeleteProduct(req model.DeleteProductRequest, resp *model.Response) error
}

type Reservation interface {
	ReserveProduct(req model.ReserveProductRequest, resp *model.Response) error
	ReleaseReserve(req model.ReleaseReserveRequest, resp *model.Response) error
}

type ProductImpl struct {
	mu   sync.RWMutex
	l    *zap.Logger
	repo *repository.Repository
}

func NewService(l *zap.Logger, repo *repository.Repository) *ProductImpl {
	return &ProductImpl{
		l:    l,
		repo: repo,
	}
}
