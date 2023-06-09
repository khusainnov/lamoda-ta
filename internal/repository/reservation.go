package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/lamoda/internal/model"
)

type ReservationRepository struct {
	db *sqlx.DB
}

func NewReservationRepository(db *sqlx.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) ReserveProduct(data model.ReserveProductRequest) error {
	query := fmt.Sprintln("INSERT INTO reservation_table (code, quantity, warehouse_id) VALUES ($1, $2, $3);")

	for _, v := range data.ProductCode {
		if _, err := r.db.Exec(query, v, data.Quantity, data.WarehouseID); err != nil {
			return fmt.Errorf("cannot reserve product, %w", err)
		}
	}

	return nil
}

func (r *ReservationRepository) ReleaseReserve(data model.ReleaseReserveRequest) error {
	query := fmt.Sprintln("UPDATE reservation_table SET departured=true WHERE code=$1 AND warehouse_id=$2;")

	for _, code := range data.ProductCode {
		if _, err := r.db.Exec(query, code, data.WarehouseID); err != nil {
			return fmt.Errorf("cannot release reserve, %w", err)
		}
	}

	return nil
}
