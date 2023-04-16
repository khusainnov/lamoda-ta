package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
)

func (p *ProductImpl) CancelReservation(req model.CancelReservationRequest, _ *model.Response) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.repo.DeleteReservation(req); err != nil {
		return fmt.Errorf("error due cancel reservation, %w", err)
	}

	return nil
}
