package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/pkg/utl/convert"
	"go-template/pkg/utl/resultwrapper"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Employee is the resolver for the employee field.
func (r *queryResolver) Employee(ctx context.Context, input gqlmodels.EmployeeQueryInput) (*gqlmodels.Employee, error) {
	employee, err := daos.FindEmployeeByID(convert.StringToInt(input.ID))
	if err != nil {
		return &gqlmodels.Employee{}, resultwrapper.ResolverSQLError(err, "data")
	}

	return convert.EmployeeToGraphQlEmployee(employee, 1), err
}

// Employees is the resolver for the employees field.
func (r *queryResolver) Employees(ctx context.Context, pagination *gqlmodels.EmployeePagination) (*gqlmodels.EmployeesPayload, error) {
	var queryMods []qm.QueryMod
	if pagination != nil {
		if pagination.Limit != 0 {
			queryMods = append(queryMods, qm.Limit(pagination.Limit), qm.Offset(pagination.Page*pagination.Limit))
		}
	}

	employees, count, err := daos.FindAllEmployeesWithCount(queryMods)
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "data")
	}
	// fmt.Printf("employees %+v\n", employees)
	// fmt.Printf("pagination %+v\n", pagination)
	return &gqlmodels.EmployeesPayload{Total: int(count), Employees: convert.EmployeesToGraphQlEmployees(employees, 1)}, nil
}
