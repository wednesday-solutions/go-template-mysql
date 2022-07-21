-- +migrate Up
CREATE TABLE incident_report_attachments
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    incident_report_id INT NOT NULL,
    url TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT ir_attachemnt_fk_incident_report_id FOREIGN KEY (incident_report_id) REFERENCES incident_reports(id)
);

-- +migrate Down
DROP TABLE incident_report_attachments;