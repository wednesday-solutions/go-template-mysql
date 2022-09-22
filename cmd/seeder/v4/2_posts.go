package main

import "go-template/cmd/seeder/utls"

func main() {

     _ = utls.SeedData("posts", `INSERT INTO posts (title, author_id) VALUES ('first post', 1);
                         INSERT INTO posts (title, author_id) VALUES ('second post', 1);
                         INSERT INTO posts (title, author_id) VALUES ('third post', 1);
                         INSERT INTO posts (title, author_id) VALUES ('fourth post', 1);
                         INSERT INTO posts (title, author_id) VALUES ('fifth post', 1);`) 
}