package mysql

import (
	"database/sql"
	"fmt"
	"go-template/pkg/utl/zaplog"
	// "go-template/testutls"
	"os"

	// DB adapter
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

// Connect ...
func Connect() (*sql.DB, error) {
	dsn := GetDSN()
	zaplog.Logger.Info("Connecting to DB\n", dsn)
	// if testutls.IsInTests() {
	// 	return sql.Open("mysql", dsn)
	// }
	return otelsql.Open("mysql", dsn, otelsql.WithAttributes(semconv.DBSystemMySQL))
}

func GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"))
	return dsn
}
