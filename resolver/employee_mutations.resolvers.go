package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/convert"
	"go-template/pkg/utl/resultwrapper"
	"go-template/pkg/utl/throttle"
	"time"

	null "github.com/volatiletech/null/v8"
)

// CreateEmployee is the resolver for the createEmployee field.
func (r *mutationResolver) CreateEmployee(ctx context.Context, input gqlmodels.EmployeeCreateInput) (*gqlmodels.Employee, error) {
	err := throttle.Check(ctx, 5, 10*time.Second)
	if err != nil {
		return nil, err
	}

	employee := models.Employee{
		Name:               input.Name,
		Email:              input.Email,
		EmployeeAccessRole: null.StringFrom(input.EmployeeAccessRole.String()),
	}

	newEmployee, err := daos.CreateEmployeeTx(employee, nil)

	graphEmployee := convert.EmployeeToGraphQlEmployee(&newEmployee, 1)

	return graphEmployee, err
}

// UpdateEmployee is the resolver for the updateEmployee field.
func (r *mutationResolver) UpdateEmployee(ctx context.Context, input gqlmodels.EmployeeUpdateInput) (*gqlmodels.Employee, error) {
	employee, err := daos.FindEmployeeByID(convert.StringToInt(input.ID))
	if err != nil {
		return nil, resultwrapper.ResolverWrapperFromMessage(404, "employee not found")
	}

	var e models.Employee
	if employee != nil {
		e = *employee
	} else {
		return nil, resultwrapper.ResolverWrapperFromMessage(404, "employee not found")
	}

	if input.Name != nil {
		e.Name = *input.Name
	}
	if input.Email != nil {
		e.Email = *input.Email
	}
	if input.EmployeeAccessRole != nil {
		e.EmployeeAccessRole = null.StringFrom(input.EmployeeAccessRole.String())
	}
	_, err = daos.UpdateEmployeeTx(e, nil)
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "new information")
	}

	graphEmployee := convert.EmployeeToGraphQlEmployee(&e, 1)

	return graphEmployee, nil
}

// DeleteEmployee is the resolver for the deleteEmployee field.
func (r *mutationResolver) DeleteEmployee(ctx context.Context, input gqlmodels.EmployeeDeleteInput) (*gqlmodels.EmployeeDeletePayload, error) {
	e, err := daos.FindEmployeeByID(convert.StringToInt(input.ID))
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "data")
	}
	_, err = daos.DeleteEmployee(*e)
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "employee")
	}
	return &gqlmodels.EmployeeDeletePayload{ID: fmt.Sprint(convert.StringToInt(input.ID))}, nil
}
