package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	todo "github.com/voslartomas/gqlgen-test/db/mongo/repositories"
	"github.com/voslartomas/gqlgen-test/graph/generated"
	"github.com/voslartomas/gqlgen-test/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := todo.Add(input)

	return todo, nil
}

func (r *mutationResolver) RemoveTodo(ctx context.Context, todoID string) (bool, error) {
	todo.Remove(todoID)

	return true, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, todoID string, data model.UpdateTodo) (*model.Todo, error) {
	todo.Update(todoID, data)

	return todo.FindByID(todoID)
}

func (r *queryResolver) Todo(ctx context.Context, todoID string) (*model.Todo, error) {
	return todo.FindByID(todoID)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todo.GetAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
