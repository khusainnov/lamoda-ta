package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
)

func (p *ProductImpl) ReserveProduct(req model.ReserveProductRequest, resp *model.Response) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.repo.ReserveProduct(req); err != nil {
		resp.Message = model.RequestErr
		return fmt.Errorf("error due reserve product, %w", err)
	}
	resp.Message = model.RequestOk

	return nil
}
