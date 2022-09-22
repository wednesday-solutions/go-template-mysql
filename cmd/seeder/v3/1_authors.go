package main

import (
     "fmt"
     "go-template/cmd/seeder/utls"
     "go-template/pkg/utl/secure"
)

func main() {

     sec := secure.New(1, nil)
     var insertQuery = fmt.Sprintf("INSERT INTO authors (first_name, last_name, username, "+
          "password, active) VALUES ('John', 'Doe', 'john', '%s', true);", sec.Hash("johndoe"))
     _ = utls.SeedData("authors", insertQuery)
}