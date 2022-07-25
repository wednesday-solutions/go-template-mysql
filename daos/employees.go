package daos

import (
	"context"
	"database/sql"
	"fmt"

	"go-template/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreateEmployeeTx ...
func CreateEmployeeTx(employee models.Employee, tx *sql.Tx) (models.Employee, error) {
	contextExecutor := getContextExecutor(tx)

	err := employee.Insert(context.Background(), contextExecutor, boil.Infer())
	return employee, err
}

// UpdateEmployeeTx ...
func UpdateEmployeeTx(employee models.Employee, tx *sql.Tx) (models.Employee, error) {
	contextExecutor := getContextExecutor(tx)
	// boil.Infer()
	fmt.Printf("boil.Infer() %+v\n", boil.Infer())
	_, err := employee.Update(context.Background(), contextExecutor, boil.Infer())
	return employee, err
}

// FindEmployeeByID ...
func FindEmployeeByID(employeeID int) (*models.Employee, error) {
	contextExecutor := getContextExecutor(nil)
	return models.FindEmployee(context.Background(), contextExecutor, employeeID)
}

// DeleteEmployee ...
func DeleteEmployee(employee models.Employee) (int64, error) {
	contextExecutor := getContextExecutor(nil)
	rowsAffected, err := employee.Delete(context.Background(), contextExecutor)
	return rowsAffected, err
}

// FindAllEmployeesWithCount ... This will get all the employees that match the queryMod filter and also return the count
func FindAllEmployeesWithCount(queryMods []qm.QueryMod) (models.EmployeeSlice, int64, error) {
	contextExecutor := getContextExecutor(nil)
	employees, err := models.Employees(queryMods...).All(context.Background(), contextExecutor)
	if err != nil {
		return models.EmployeeSlice{}, 0, err
	}
	queryMods = append(queryMods, qm.Offset(0))
	count, err := models.Employees(queryMods...).Count(context.Background(), contextExecutor)
	return employees, count, err
}
