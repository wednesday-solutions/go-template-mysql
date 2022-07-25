package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
)

// CaseStatus is the resolver for the caseStatus field.
func (r *queryResolver) CaseStatus(ctx context.Context, input gqlmodels.CaseStatusQueryInput) (*gqlmodels.CaseStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

// CaseStatuses is the resolver for the caseStatuses field.
func (r *queryResolver) CaseStatuses(ctx context.Context, pagination *gqlmodels.CaseStatusPagination) (*gqlmodels.CaseStatusesPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns gqlmodels.QueryResolver implementation.
func (r *Resolver) Query() gqlmodels.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
