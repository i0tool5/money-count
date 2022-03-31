
CREATE TABLE payments 
(
    id          bigserial PRIMARY KEY, 
    user_id     bigint not null, 
    date date   DEFAULT current_date,
    type        text,
    description text,
    amount      bigint
); 
