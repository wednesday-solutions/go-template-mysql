-- +migrate Up
CREATE TABLE articles (
				id INT AUTO_INCREMENT PRIMARY KEY,
				title TEXT,
				author_id INT,
				created_at TIMESTAMP DEFAULT NOW(),
				updated_at TIMESTAMP,
				deleted_at TIMESTAMP,
				CONSTRAINT fk__author_articles FOREIGN KEY (author_id) REFERENCES authors(id)
			);
CREATE INDEX author_id_idx ON articles(author_id);

-- +migrate Down
DROP TABLE articles;