package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/cnvrttogql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id int) (*gqlmodels.Author, error) {
	var contextExecutor = boil.GetContextDB()

	author, err := models.FindAuthor(ctx, contextExecutor, id)
	if err != nil {
		return nil, err
	}
	return cnvrttogql.AuthorToGraphQLAuthor(author), nil
}

// AllAuthors is the resolver for the allAuthors field.
func (r *queryResolver) AllAuthors(ctx context.Context) ([]*gqlmodels.Author, error) {
	var contextExecutor = boil.GetContextDB()
	authors, err := models.Authors().All(ctx, contextExecutor)
	if err != nil {
		return nil, err
	}
	return cnvrttogql.AuthorsToGraphQLAuthors(authors), nil
}
