CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE warehouse_table (
    id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    available BOOL DEFAULT FALSE
);

CREATE TABLE product_table(
    id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    size VARCHAR NOT NULL,
    code VARCHAR NOT NULL,
    quantity INT DEFAULT 0 CHECK (quantity >= 0),
    warehouse_id UUID,
    FOREIGN KEY (warehouse_id)
        REFERENCES warehouse_table(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE reservation_table (
    id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    code VARCHAR NOT NULL,
    quantity INT DEFAULT 1 CHECK (quantity >= 1),
    warehouse_id UUID,
    FOREIGN KEY (warehouse_id)
        REFERENCES warehouse_table(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE OR REPLACE FUNCTION fn_reduce_product_quantity()
    RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$$
    BEGIN
        UPDATE product_table SET quantity = quantity - new.quantity WHERE code = new.code AND warehouse_id = new.warehouse_id;
        RETURN new;
    END;
$$;

CREATE OR REPLACE TRIGGER tr_insert_reservation AFTER INSERT ON reservation_table
    FOR EACH ROW EXECUTE FUNCTION fn_reduce_product_quantity();

insert into warehouse_table("name", available) values ('Warehouse 1', true);
insert into warehouse_table("name", available) values ('Warehouse 2', false);
insert into warehouse_table("name", available) values ('Warehouse 3', true);

CREATE OR REPLACE FUNCTION fn_rollback_product()
    RETURNS TRIGGER
    LANGUAGE plpgsql
AS
    $$
    BEGIN
        UPDATE product_table SET quantity = quantity + old.quantity WHERE code = old.code AND warehouse_id = old.warehouse_id;
        RETURN old;
    END;
$$;

CREATE OR REPLACE TRIGGER tr_delete_reservation BEFORE DELETE ON reservation_table
    FOR EACH ROW EXECUTE FUNCTION fn_rollback_product();

