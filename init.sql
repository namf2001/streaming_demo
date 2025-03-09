CREATE TABLE records
(
    id        SERIAL PRIMARY KEY,
    keyword   VARCHAR(255),
    node      VARCHAR(255),
    create_at DATE,
    update_at DATE
);

INSERT INTO records (keyword, node, create_at, update_at)
VALUES ('key1', 'node1', '2023-01-01', '2023-01-01'),
       ('key2', 'node2', '2023-01-02', '2023-01-02'),
       ('key3', 'node3', '2023-01-03', '2023-01-03');