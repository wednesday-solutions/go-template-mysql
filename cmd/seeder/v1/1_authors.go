package main

import (
	"fmt"
	"go-template/cmd/seeder/utls"
	"go-template/pkg/utl/secure"
)

func main() {

	sec := secure.New(1, nil)
	insertQuery := fmt.Sprintf("INSERT INTO authors(id, first_name, last_name, username, password, active)"+
					"VALUES (1, 'john', 'doe', 'johndoe', '%s', true);"+
				"INSERT INTO authors(id, first_name, last_name, username, password, active)"+
					"VALUES (2, 'author', '2', 'author2', '%s', true);"+
				"INSERT INTO authors(id, first_name, last_name, username, password, active)"+
					"VALUES (3, 'author', '3', 'author3', '%s', true);", sec.Hash("johndoe"), sec.Hash("author"), sec.Hash("author"))

	_ = utls.SeedData("authors", insertQuery)
}