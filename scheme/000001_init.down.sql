DROP TABLE warehouse_table CASCADE;
DROP TABLE product_table;
DROP TABLE reservation_table;
DROP FUNCTION IF EXISTS fn_reduce_product_quantity();
DROP TRIGGER IF EXISTS tr_insert_reservastion ON reservation_table;
DROP FUNCTION IF EXISTS fn_release_product();
DROP TRIGGER IF EXISTS tr_release_product ON reservation_table;
