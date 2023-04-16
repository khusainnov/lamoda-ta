package productservice

import (
	"fmt"

	"github.com/khusainnov/lamoda/internal/model"
	"go.uber.org/zap"
)

func (p *ProductImpl) UploadProduct(req model.UploadProductRequest, resp *model.UploadProductResponse) error {
	p.l.Info("Product", zap.Any("Name", req.Name))

	if err := p.repo.UploadProduct(req); err != nil {
		return fmt.Errorf("cannot upload product, %w", err)
	}

	resp.Message = model.RequestOk

	return nil
}
