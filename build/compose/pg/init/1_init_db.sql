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

CREATE TABLE IF NOT EXISTS goods(
    id BIGINT NOT NULL,
    name varchar(100) NOT NULL,
    size varchar(100) NOT NULL DEFAULT '1x1x1'::character varying,
    quantity BIGINT NOT NULL,
    reserve jsonb NOT NULL DEFAULT '[]'::jsonb,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS warehouses(
    id SERIAL NOT NULL,
    status boolean,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS reserves(
    id varchar(50) NOT NULL,
    goods_id bigint NOT NULL,
    wh_id bigint NOT NULL,
    quantity bigint NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_reserves_goods FOREIGN key(goods_id) REFERENCES goods(id),
    CONSTRAINT fk_reserves_warehouses FOREIGN key(wh_id) REFERENCES warehouses(id)
);

