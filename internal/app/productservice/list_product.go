package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
)

func (p *ProductImpl) ListProduct(req model.ListProductRequest, resp *[]model.Product) error {
	rsp, err := p.repo.ListProduct(req)
	if err != nil {
		return fmt.Errorf("cannot get product list, %w", err)
	}

	*resp = *rsp

	return nil
}
