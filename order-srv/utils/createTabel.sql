CREATE TABLE orders (
    id SERIAL NOT NULL PRIMARY KEY,
    sid INTEGER NOT NULL, 
    name VARCHAR(50) NOT NULL DEFAULT '',
    create_time TIMESTAMP NOT NULL DEFAULT now()
);

COMMENT ON COLUMN orders.id IS '主键';
COMMENT ON COLUMN orders.sid IS '库存ID';
COMMENT ON COLUMN orders.name IS '商品名称';
COMMENT ON COLUMN stock.create_time IS '创建时间';