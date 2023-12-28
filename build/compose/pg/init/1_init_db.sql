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

INSERT INTO warehouses (status) VALUES
    (true),
    (true),
    (false);

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

INSERT INTO goods (name, size, quantity, wh_id) VALUES 
    ('cap', '5x5x6',40,1),
    ('kettle', '10x10x10',10, 2),
    ('mug', '6x6x7', 15, 1),
    ('ewewew','4x8x8', 5, 1);




