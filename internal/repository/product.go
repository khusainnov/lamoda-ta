package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/lamoda/internal/model"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) UploadProduct(product model.UploadProductRequest) error {
	query := fmt.Sprintln("INSERT INTO product_table (name, size, code, quantity, warehouse_id) VALUES ($1, $2, $3, $4, $5);")

	if _, err := p.db.Exec(query, product.Name, product.Size, product.Code, product.Quantity, product.WarehouseID); err != nil {
		return fmt.Errorf("cannot insert data, %w", err)
	}

	return nil
}

func (p *ProductRepository) UpdateProduct(data model.UpdateProductRequest) error {
	query := fmt.Sprintln("UPDATE product_table SET name=$1, size=$2, code=$3, quantity=$4, warehouse_id=$5 WHERE code=$3 and warehouse_id=$5;")

	if _, err := p.db.Exec(query, data.Name, data.Size, data.Code, data.Quantity, data.WarehouseID); err != nil {
		return fmt.Errorf("error due update product, %w", err)
	}

	return nil
}

func (p *ProductRepository) GetProduct(data model.GetProductRequest) (*model.Product, error) {
	var resp model.Product
	query := fmt.Sprintln("SELECT * FROM product_table WHERE code=$1 AND warehouse_id=$2;")

	if err := p.db.Get(&resp, query, data.ProductCode, data.WarehouseID); err != nil {
		return nil, fmt.Errorf("error due select product, %w", err)
	}

	return &resp, nil
}

func (p *ProductRepository) ListProduct(data model.ListProductRequest) (*[]model.Product, error) {
	var resp []model.Product

	query := fmt.Sprintln("SELECT * FROM product_table WHERE warehouse_id=$1;")

	if err := p.db.Select(&resp, query, data.WarehouseID); err != nil {
		return nil, fmt.Errorf("error due select product list, %w", err)
	}

	return &resp, nil
}

func (p *ProductRepository) DeleteProduct(data model.DeleteProductRequest) error {
	query := fmt.Sprintln("DELETE FROM product_table WHERE code=$1 and warehouse_id=$2;")

	if err := p.db.QueryRow(query, data.ProductCode, data.WarehouseID).Err(); err != nil {
		return fmt.Errorf("cannot delete product, %w", err)
	}

	return nil
}
