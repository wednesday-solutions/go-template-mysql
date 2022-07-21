-- +migrate Up
CREATE TABLE incident_reports
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    employee_id INT NOT NULL,
    case_manager_id INT,
    line_manager_id INT,
    core_system_rider_id INT NOT NULL,
    incident_details TEXT NOT NULL,
    incident_date TIMESTAMP NOT NULL,
    evidence_receipt_date TIMESTAMP,
    remarks TEXT,
    type ENUM ('MINOR', 'MAJOR'),
    url TEXT,
    preventive_suspension BOOLEAN,
    date_of_administrative_hearing TIMESTAMP,
    case_folder TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT ir_fk_employee_id FOREIGN KEY (employee_id) REFERENCES employees(id),
    CONSTRAINT ir_fk_case_manager_id FOREIGN KEY (case_manager_id) REFERENCES employees(id),
    CONSTRAINT ir_fk_line_manager_id FOREIGN KEY (line_manager_id) REFERENCES employees(id)
);

-- +migrate Down
DROP TABLE incident_reports;