CREATE TABLE stock(
    id SERIAL NOT NULL PRIMARY KEY COMMENT '主键',
    name VARCHAR(50) NOT NULL,
    count INTEGER NOT NULL,
    sale INTEGER NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT now()
);

COMMENT ON COLUMN stock.id IS '主键';;
COMMENT ON COLUMN stock.name IS '名称';;
COMMENT ON COLUMN stock.count IS '库存';
COMMENT ON COLUMN stock.sale IS '销量';
COMMENT ON COLUMN stock.create_time IS '创建时间';