BEGIN;
  CREATE TABLE products (
    id SERIAL
    --
    ,name   text unique
    ,description  varchar(255)
    ,price  numeric
    ,image_url  varchar(255)
    --
    ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
    ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
    -- soft delete
    ,deleted_at   timestamptz
    --
    ,PRIMARY KEY (id)
  );

  CREATE INDEX "idx_products_deleted_at" ON "products" ("deleted_at");
  CREATE INDEX "idx_products_name" ON "products" ("name");
COMMIT;
