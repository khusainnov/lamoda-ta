package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
)

func (p *ProductImpl) ReleaseReserve(req model.ReleaseReserveRequest, resp *model.Response) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.repo.ReleaseReserve(req); err != nil {
		resp.Message = model.RequestErr
		return fmt.Errorf("error due cancel reservation, %w", err)
	}

	resp.Message = model.RequestOk

	return nil
}
