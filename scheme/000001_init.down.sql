DROP TABLE product_table;
DROP TABLE warehouse_table;
DROP TABLE reservation_table;
DROP FUNCTION IF EXISTS fn_reduce_product_quantity();
DROP TRIGGER IF EXISTS tr_insert_reservation ON reservation_table;