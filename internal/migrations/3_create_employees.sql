-- +migrate Up
CREATE TABLE employees
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL,
    employee_access_role ENUM ('EMPLOYEE', 'CASE_MANAGER', 'LINE_MANAGER', 'SUPER') DEFAULT 'EMPLOYEE',

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE INDEX name_idx ON employees(name);
CREATE INDEX email_idx ON employees(email);
CREATE INDEX employee_access_role_idx ON employees(employee_access_role);

-- +migrate Down
DROP TABLE employees;