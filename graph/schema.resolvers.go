package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/uch-kuduk/naive-graphql.git/graph/generated"
	"github.com/uch-kuduk/naive-graphql.git/graph/model"
)

func (r *mutationResolver) CreateCar(ctx context.Context, input model.CreateCarInput) (*model.CreateCarPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCar(ctx context.Context, input model.DeleteCarInput) (*model.DeleteCarPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CarBrands(ctx context.Context) (*model.CarBrandsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserCars(ctx context.Context) (*model.CarsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserCar(ctx context.Context, id string) (*model.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

/* TO DELETE
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
*/
