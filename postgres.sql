-- This is postgres
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    position VARCHAR(255),
    salary FLOAT
);