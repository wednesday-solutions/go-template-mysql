package main

import (
	"fmt"

	"go-template/cmd/seeder/utls"
	// "go-template/models"
	// "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {

	// getting the latest location company and role id so that we can seed a new user
	// role, _ := models.Roles(qm.OrderBy("id ASC")).One(context.Background(), db)

	var insertQuery = fmt.Sprintf("INSERT INTO articles (id, title, author_id) VALUES (1, 'first article', 1);"+
		"INSERT INTO articles (id, title, author_id) VALUES (2, 'second article', 1);"+
		"INSERT INTO articles (id, title, author_id) VALUES (3, 'third article', 1);"+
		"INSERT INTO articles (id, title, author_id) VALUES (4, 'fourth article', 2);"+
		"INSERT INTO articles (id, title, author_id) VALUES (5, 'fifth article', 2);"+
		"INSERT INTO articles (id, title, author_id) VALUES (6, 'sixth article', 3);")
		
	_ = utls.SeedData("users", insertQuery)
}