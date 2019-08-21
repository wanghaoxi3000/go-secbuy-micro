CREATE TABLE payments (
    id SERIAL NOT NULL PRIMARY KEY,
    sid INTEGER NOT NULL,
    name VARCHAR(50) NOT NULL DEFAULT '',
    state INTEGER NOT NULL DEFAULT 1,
    create_time TIMESTAMP NOT NULL DEFAULT now()
);

COMMENT ON COLUMN payment.id IS '主键';
COMMENT ON COLUMN payment.sid IS '订单ID';
COMMENT ON COLUMN payment.state IS '订单状态';
COMMENT ON COLUMN payment.name IS '商品名称';
COMMENT ON COLUMN payment.create_time IS '创建时间';