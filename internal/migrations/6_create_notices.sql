-- +migrate Up
CREATE TABLE notices
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    incident_report_id INT NOT NULL,
    type ENUM ('NOTICE_OF_WARNING', 'NOTICE_TO_EXPLAIN', 'NOTICE_OF_DECISION') NOT NULL,
    issue_date TIMESTAMP,
    url TEXT,
    sla INT,
    notice_read BOOLEAN,
    notice_signed BOOLEAN,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT notice_fk_incident_report_id FOREIGN KEY (incident_report_id) REFERENCES incident_reports(id)
);

-- +migrate Down
DROP TABLE notices;