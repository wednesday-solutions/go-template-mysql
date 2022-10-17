package main

import (
	"context"
	"fmt"

	"go-template/cmd/seeder/utls"
	"go-template/internal/mysql"
	"go-template/models"
	"go-template/pkg/utl/secure"
	"go-template/pkg/utl/zaplog"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {

	sec := secure.New(1, nil)
	db, _ := mysql.Connect()
	role, err := models.Roles(qm.OrderBy("id ASC")).One(context.Background(), db)
	if err != nil {
		zaplog.Logger.Error("error while seeding", err)
	}
	var insertQuery = fmt.Sprintf("INSERT INTO users (first_name, last_name, username, password, "+
		"email, active, role_id) VALUES ('Mohammed Ali', 'Chherawalla', 'admin', '%s', 'johndoe@mail.com', true, %d);",
		sec.Hash("adminuser"), role.ID)
	_ = utls.SeedData("users", insertQuery)
}
