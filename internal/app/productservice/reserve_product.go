package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
)

func (p *ProductImpl) ReserveProduct(req model.ReserveProductRequest, _ *model.Response) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.repo.ReserveProduct(req); err != nil {
		return fmt.Errorf("error due reserve product, %w", err)
	}

	return nil
}
