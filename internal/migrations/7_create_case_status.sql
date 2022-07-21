-- +migrate Up
CREATE TABLE case_status
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    incident_report_id INT NOT NULL,
    status ENUM ('FILED', 
                'CASE_MANAGER_ASSIGNED', 
                'VALID_BY_CASE_MANAGER', 
                'INCOMPLETE_EVIDENCE', 
                'NO_DISCIPLINARY_CASE', 
                'LINE_MANAGER_ASSIGNED',
                'VALID_BY_LINE_MANAGER',
                'NTE_DRAFTED',
                'NTE_ISSUED',
                'RIDER_RESPONSE',
                'ADMIN_HEARING',
                'NOD_DRAFTED',
                'NOD_ISSUED',
                'NOW_DRAFTED',
                'NOW_ISSUED') NOT NULL,
    status_remarks TEXT NOT NULL,
    user_id INT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT case_status_fk_incident_report_id FOREIGN KEY (incident_report_id) REFERENCES incident_reports(id),
    CONSTRAINT case_status_fk_user_id FOREIGN KEY (user_id) REFERENCES employees(id)
);

-- +migrate Down
DROP TABLE case_status;