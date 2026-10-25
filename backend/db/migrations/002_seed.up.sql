INSERT INTO products (sku, name, price, stock, is_active)
VALUES
    ('SPT-BLUE-F40', 'Flatshoes', 220000, 27, true),
    ('SPT-PINK-F39', 'Flatshoes', 215000, 18, true),
    ('SPT-BLUE-F39', 'Flatshoes', 215000, 12, true),
    ('SPT-GREN-F40', 'Flatshoes', 199000, 76, true),
    ('BLS-RED-XL', 'Blouse Lengan Balon', 275000, 11, true),
    ('BLS-BRWN-S', 'Blouse Lengan Balon', 265000, 0, false),
    ('BLS-RED-M', 'Blouse Lengan Balon', 215000, 8, false),
    ('KLT-PINK-XS', 'Kulot Plisket', 129000, 19, true),
    ('KLT-WHTE-XL', 'Kulot Plisket', 129000, 213, true),
    ('KLT-BLUE-S', 'Kulot Plisket', 129000, 41, true),
    ('KLT-GRAY-M', 'Kulot Plisket', 129000, 68, false),
    ('JKT-WHTE-2XL', 'Jaket Bomber', 215000, 11, true),
    ('JKT-BLCK-L', 'Jaket Bomber', 319000, 24, true),
    ('JKT-GRAY-S', 'Jaket Bomber', 325000, 125, false),
    ('JKT-BLUE-XL', 'Jaket Bomber', 312000, 27, false),
    ('JKT-BLCK-M', 'Jaket Bomber', 318000, 133, false),
    ('JKT-BLUE-L', 'Jaket Bomber', 325000, 34, true);

INSERT INTO sales_orders (customer, status, total, created_at, cancelled_at)
VALUES
    ('Dina', 'CREATED', 215000, '2026-08-15 19:36:00+07', NULL),
    ('Budi', 'CREATED', 1807000, '2026-08-17 17:05:21+07', NULL),
    ('Siti', 'CREATED', 404000, '2026-08-22 13:47:10+07', NULL),
    ('Sarah', 'CANCELLED', 220000, '2026-08-23 00:16:32+07', '2026-08-24 08:12:16+07');

INSERT INTO sales_order_items (order_id, sku, qty, unit_price)
VALUES
    (1, 'SPT-PINK-F39', 1, 215000),
    (2, 'JKT-BLCK-L', 2, 319000),
    (2, 'JKT-WHTE-2XL', 1, 215000),
    (2, 'JKT-BLCK-M', 3, 318000),
    (3, 'BLS-RED-XL', 1, 275000),
    (3, 'KLT-GRAY-M', 1, 129000),
    (4, 'SPT-BLUE-F40', 1, 220000);