CREATE TABLE payments (
    id bigserial PRIMARY KEY,
    date date DEFAULT current_date,
    type text,
    description text,
    amount bigint
);
