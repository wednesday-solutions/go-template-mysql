package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/pkg/utl/convert"
	"go-template/pkg/utl/resultwrapper"
)

// IncidentReport is the resolver for the incidentReport field.
func (r *queryResolver) IncidentReport(ctx context.Context, input gqlmodels.IncidentReportQueryInput) (*gqlmodels.IncidentReport, error) {
	incidentReport, err := daos.FindIncidentReportByID(convert.StringToInt(input.ID))
	if err != nil {
		return &gqlmodels.IncidentReport{}, resultwrapper.ResolverSQLError(err, "data")
	}

	return convert.IncidentReportToGraphQlIncidentReport(incidentReport, 1), err
}

// IncidentReports is the resolver for the incidentReports field.
func (r *queryResolver) IncidentReports(ctx context.Context, pagination *gqlmodels.IncidentReportPagination) (*gqlmodels.IncidentReportsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}
