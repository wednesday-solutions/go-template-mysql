package resolver_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	fm "go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/convert"
	"go-template/resolver"
	"go-template/testutls"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/stretchr/testify/assert"
)

func TestEmployees(t *testing.T) {
	cases := []struct {
		name       string
		pagination *fm.EmployeePagination
		wantResp   []*models.Employee
		wantErr    bool
	}{
		{
			name:     "Success",
			wantErr:  false,
			wantResp: testutls.MockEmployees(),
		},
	}

	resolver1 := resolver.Resolver{}
	for _, tt := range cases {
		t.Run(
			tt.name,
			func(t *testing.T) {
				err := godotenv.Load("../.env.local")
				if err != nil {
					fmt.Print("error loading .env file")
				}
				db, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				// Injectmockinstanceintoboil.
				oldDB := boil.GetDB()
				defer func() {
					db.Close()
					boil.SetDB(oldDB)
				}()
				boil.SetDB(db)

				// get employee by id
				rows := sqlmock.
					NewRows(
						[]string{"id", "email", "name", "employee_access_role"},
					).
					AddRow(testutls.MockID, testutls.MockEmail, "First", "CASE_MANAGER")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT `employees`.* FROM `employees`;")).
					WithArgs().
					WillReturnRows(rows)
				rowCount := sqlmock.NewRows([]string{"count"}).
					AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM `employees`;")).
					WithArgs().
					WillReturnRows(rowCount)

				c := context.Background()
				ctx := context.WithValue(c, testutls.EmployeeKey, testutls.MockEmployee())
				response, err := resolver1.Query().Employees(ctx, tt.pagination)

				if tt.wantResp != nil &&
					response != nil {
					assert.Equal(t, len(tt.wantResp), len(response.Employees))
				}
				assert.Equal(t, tt.wantErr, err != nil)
			},
		)
	}
}

func TestEmployee(
	t *testing.T,
) {
	type args struct {
		input fm.EmployeeQueryInput
	}
	cases := []struct {
		name     string
		wantResp *fm.Employee
		wantErr  bool
		args     args
	}{
		{
			name: "Success",
			args: args{
				input: testutls.MockEmployeeQueryInput(),
			},
			wantResp: convert.EmployeeToGraphQlEmployee(
				testutls.MockEmployee(),
				4,
			),
		},
	}

	_, db, err := testutls.SetupEnvAndDB(
		t,
		testutls.Parameters{
			EnvFileLocation: `../.env.local`,
		},
	)
	if err != nil {
		panic(
			"failed to setup env and db",
		)
	}
	oldDb := boil.GetDB()
	boil.SetDB(db)
	defer func() {
		db.Close()
		boil.SetDB(
			oldDb,
		)
	}()
	resolver1 := resolver.Resolver{}
	for _, tt := range cases {
		t.Run(
			tt.name,
			func(t *testing.T) {
				err := godotenv.Load("../.env.local")
				if err != nil {
					fmt.Print("error loading .env file")
				}
				db, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				// Injectmockinstanceintoboil.
				oldDB := boil.GetDB()
				defer func() {
					db.Close()
					boil.SetDB(oldDB)
				}()
				boil.SetDB(db)

				t.Logf("LOG CHECK\n")

				// get employee by id
				rows := sqlmock.
					NewRows(
						[]string{"id", "email", "name", "employee_access_role"},
					).
					AddRow(testutls.MockID, testutls.MockEmail, "First", "CASE_MANAGER")

				t.Logf("rows %+v\n", rows)

				mock.ExpectQuery(regexp.QuoteMeta("select * from `employees` where `id`=?")).
					WithArgs(testutls.MockID).
					WillReturnRows(rows)
				// rowCount := sqlmock.NewRows([]string{"count"}).
				//     AddRow(1)
				// mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM `employees`;")).
				//     WithArgs().
				//     WillReturnRows(rowCount)

				c := context.Background()
				ctx := context.WithValue(c, testutls.EmployeeKey, testutls.MockEmployee())
				response, err := resolver1.Query().Employee(ctx, tt.args.input)
				t.Logf("response %+v\n", response)
				if tt.wantResp != nil &&
					response != nil {
					assert.Equal(t, tt.wantResp, response)
				}
				assert.Equal(t, tt.wantErr, err != nil)

				// b, _ := json.Marshal(tt.args.employee)
				// c := context.Background()
				// ctx := context.WithValue(
				// 	c,
				// 	testutls.EmployeeKey,
				// 	testutls.MockEmployee(),
				// )
				// response, _ := resolver1.Query().Me(ctx)
				// assert.Equal(t, tt.wantResp, response)
			},
		)
	}
}
