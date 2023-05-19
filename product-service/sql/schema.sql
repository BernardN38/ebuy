CREATE TABLE products
(
    id         serial PRIMARY KEY,
    user_id int NOT NULL,
    name  text NOT NULL,
    description text NOT NULL,
    price int NOT NULL,
    upload_date timestamp NOT NULL DEFAULT NOW()
);