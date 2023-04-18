package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
)

func (p *ProductImpl) DeleteProduct(req model.DeleteProductRequest, resp *model.Response) error {
	fmt.Printf("\n%v\n", req)

	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.repo.DeleteProduct(req); err != nil {
		resp.Message = model.RequestErr
		return fmt.Errorf("error due deleting product")
	}
	resp.Message = model.RequestOk

	return nil
}
