SELECT 'CREATE DATABASE zadanie' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'zadanie')\gexec

DO 
$$ 
BEGIN 
    IF NOT EXISTS (SELECT * FROM pg_user WHERE usename = 'backend_user') THEN	
        CREATE ROLE backend_user password 'secret_password';
    END IF;
END 
$$ 
;

GRANT ALL PRIVILEGES ON DATABASE zadanie TO backend_user;

DROP TABLE IF EXISTS warehouses;
CREATE TABLE warehouses(
    id SERIAL NOT NULL,
    status boolean,
    PRIMARY KEY(id)
);

DROP TABLE IF EXISTS goods;
CREATE TABLE goods(
    id SERIAL NOT NULL,
    name varchar(100) NOT NULL,
    size varchar(100) NOT NULL DEFAULT '1x1x1'::character varying,
    quantity bigint,
    wh_id bigint NOT NULL,
    reserve_id varchar(50),
    PRIMARY KEY(id),
    CONSTRAINT fk_goods_warehouses FOREIGN key(wh_id) REFERENCES warehouses(id)
);




DROP TABLE IF EXISTS stocks;
-- CREATE TABLE IF NOT EXISTS stocks(
--     id SERIAL NOT NULL,
--     goods_id bigint NOT NULL,
--     wh_id bigint NOT NULL,
--     quantity bigint NOT NULL,
--     PRIMARY KEY(id),
--     CONSTRAINT fk_stocks_goods FOREIGN key(goods_id) REFERENCES goods(id),
--     CONSTRAINT fk_stocks_warehouses FOREIGN key(wh_id) REFERENCES warehouses(id)
-- );

DROP TABLE IF EXISTS reserves;
-- CREATE TABLE IF NOT EXISTS reserves(
--     id varchar(50) NOT NULL,
--     stock_id bigint NOT NULL,
--     quantity bigint NOT NULL,
--     PRIMARY KEY(id),
--     CONSTRAINT fk_reserves_stock FOREIGN key(stock_id) REFERENCES stocks(id)
-- );


