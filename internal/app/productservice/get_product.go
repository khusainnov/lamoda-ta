package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
)

func (p *ProductImpl) GetProduct(req model.GetProductRequest, resp *model.Product) error {
	rsp, err := p.repo.GetProduct(req)
	if err != nil {
		return fmt.Errorf("cannot get product, %w", err)
	}

	*resp = *rsp

	return nil
}
