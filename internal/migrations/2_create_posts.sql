-- +migrate Up
CREATE TABLE posts (
          id INT AUTO_INCREMENT PRIMARY KEY,
          title TEXT,
          author_id INT,
          created_at TIMESTAMP,
          updated_at TIMESTAMP,
          deleted_at TIMESTAMP,
          CONSTRAINT fk_author_posts FOREIGN KEY (author_id) REFERENCES authors(id)
);

-- +migrate Down
DROP CONSTRAINT fk_author_posts;
DROP TABLE posts;