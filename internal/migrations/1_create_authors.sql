-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE authors (
			id INT AUTO_INCREMENT PRIMARY KEY,
			first_name TEXT,
			last_name TEXT,
			username VARCHAR(100) UNIQUE,
			password TEXT,
			active BOOLEAN,
			token TEXT,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP
		);
CREATE INDEX username_idx ON authors(username);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE authors;