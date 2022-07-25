package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
)

// Notice is the resolver for the notice field.
func (r *queryResolver) Notice(ctx context.Context, input gqlmodels.NoticeQueryInput) (*gqlmodels.Notice, error) {
	panic(fmt.Errorf("not implemented"))
}

// Notices is the resolver for the notices field.
func (r *queryResolver) Notices(ctx context.Context, pagination *gqlmodels.NoticePagination) (*gqlmodels.NoticesPayload, error) {
	panic(fmt.Errorf("not implemented"))
}
