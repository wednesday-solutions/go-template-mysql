-- +migrate Up
CREATE TABLE posts (
				id INT AUTO_INCREMENT PRIMARY KEY,
				title TEXT,
				author_id int,
				created_at TIMESTAMP DEFAULT NOW(),
				updated_at TIMESTAMP,
				deleted_at TIMESTAMP,
				CONSTRAINT fk__author_posts FOREIGN KEY (author_id) REFERENCES authors(id)
			);

-- +migrate Down
DROP CONSTRAINT fk__author_posts;
DROP TABLE posts;