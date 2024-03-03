package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"author/graph/model"
	"context"
)

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	authors := make([]*model.Author, 3)
	name1, name2, name3 := "Agatha Cristie", "J.K.Rowling", "H.G.Well"
	age1, age2, age3 := 98, 54, 78
	authors[0] = &model.Author{
		ID:   1,
		Name: &name1,
		Age:  &age1,
	}
	authors[1] = &model.Author{
		ID:   2,
		Name: &name2,
		Age:  &age2,
	}
	authors[2] = &model.Author{
		ID:   3,
		Name: &name3,
		Age:  &age3,
	}
	return authors, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
