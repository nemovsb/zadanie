INSERT INTO goods (name, size) VALUES 
    ('cap', '5x5x6'),
    ('kettle', '10x10x10'),
    ('mug', '6x6x7'),
    ('ewewew','4x8x8');
    
INSERT INTO warehouses (status) VALUES
    (true),
    (true),
    (false);


INSERT INTO stocks (goods_id, wh_id, quantity) 
    VALUES  (1, 1, 5),
            (2, 1, 1),
            (3, 2, 10);

INSERT INTO reserves (id, stock_id, quantity) VALUES
    ('test_reserve_1', 1, 1),
    ('test_reserve_2', 1, 1);
