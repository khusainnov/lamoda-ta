package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
)

func (p *ProductImpl) UpdateProduct(req model.UpdateProductRequest, resp *model.Response) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.repo.UpdateProduct(req); err != nil {
		return fmt.Errorf("cannot update product, %w", err)
	}

	return nil
}
