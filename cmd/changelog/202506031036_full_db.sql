-- +goose Up
-- +goose StatementBegin

INSERT INTO products (uuid, store_cost, store_amount) VALUES
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 120.50, 100),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 340.00, 50),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 75.75, 200),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 480.20, 30),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 1500.00, 10),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 600.00, 40),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a17', 220.10, 80),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 55.30, 120),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 90.00, 70),
                                                          ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a20', 130.00, 60);

INSERT INTO supply (uuid, comment, creation_date, desired_date, status, responsible_user, edited, edited_date, cost) VALUES
                                                                                                                         ('d1a3c541-6c6b-4a6e-9c6b-111111111111', 'Поставка №1 - электроника', '2025-05-01 10:00:00', '2025-05-10 12:00:00', 'in_work', 'user_1234567890123456', false, NULL, 120000.00),
                                                                                                                         ('d1a3c541-6c6b-4a6e-9c6b-222222222222', 'Поставка №2 - офисные товары', '2025-05-03 11:30:00', '2025-05-15 18:00:00', 'created', 'user_2345678901234567', false, NULL, 54000.00),
                                                                                                                         ('d1a3c541-6c6b-4a6e-9c6b-333333333333', 'Поставка №3 - бытовая техника', '2025-05-05 09:00:00', '2025-05-12 15:00:00', 'on_the_road', 'user_3456789012345678', true, '2025-05-06 14:00:00', 98000.00),
                                                                                                                         ('d1a3c541-6c6b-4a6e-9c6b-444444444444', 'Поставка №4 - промтовары', '2025-05-07 08:45:00', '2025-05-20 16:00:00', 'done', 'user_4567890123456789', false, NULL, 72000.00);

INSERT INTO supply_product (product_uuid, supply_uuid, amount) VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd1a3c541-6c6b-4a6e-9c6b-111111111111', 50),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'd1a3c541-6c6b-4a6e-9c6b-111111111111', 30),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'd1a3c541-6c6b-4a6e-9c6b-222222222222', 80),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'd1a3c541-6c6b-4a6e-9c6b-222222222222', 15),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'd1a3c541-6c6b-4a6e-9c6b-333333333333', 5),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 'd1a3c541-6c6b-4a6e-9c6b-333333333333', 10),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a17', 'd1a3c541-6c6b-4a6e-9c6b-444444444444', 20),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 'd1a3c541-6c6b-4a6e-9c6b-444444444444', 40);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

delete from supply_product where id > 0;
delete from product where id > 0;
delete from supply where id > 0;

-- +goose StatementEnd