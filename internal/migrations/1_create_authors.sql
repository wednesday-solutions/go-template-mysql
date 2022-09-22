-- +migrate Up
CREATE TABLE authors (
          id INT AUTO_INCREMENT PRIMARY KEY,
          first_name TEXT,
          last_name TEXT,
          username VARCHAR(100) UNIQUE,
          password TEXT,
          active BOOLEAN,
          token TEXT,
          creted_at TIMESTAMP,
          updated_at TIMESTAMP,
          deleted_at TIMESTAMP
     );
CREATE INDEX username_idx ON authors(username);

-- +migrate Down
DROP TABLE authors;