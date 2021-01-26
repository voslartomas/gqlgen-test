package todo

import (
	"context"
	"fmt"

	mongodb "github.com/voslartomas/gqlgen-test/db/mongo"
	"github.com/voslartomas/gqlgen-test/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getTodoRepository() *mongo.Collection {
	return mongodb.GetDatabase().Collection("todos")
}

func Add(input model.NewTodo) *model.Todo {
	todo := &model.Todo{
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "user" + input.UserID,
		},
	}

	insertResult, err := getTodoRepository().InsertOne(context.TODO(), todo)
	fmt.Println(insertResult)

	if err != nil {
		panic(err)
	}

	return todo
}

func GetAll() []*model.Todo {
	cursor, err := getTodoRepository().Find(context.TODO(), bson.D{{}})

	if err != nil {
		panic(err)
	}

	var todos []*model.Todo
	if err = cursor.All(context.TODO(), &todos); err != nil {
		panic(err)
	}

	fmt.Println(todos)

	return todos
}

func Remove(todoID string) bool {
	id, _ := primitive.ObjectIDFromHex(todoID)
	result, err := getTodoRepository().DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted todo with id "+todoID, 1 == result.DeletedCount)

	return true
}

func Update(todoID string, data model.UpdateTodo) (*model.Todo, error) {
	id, _ := primitive.ObjectIDFromHex(todoID)
	result, err := getTodoRepository().UpdateOne(context.TODO(), bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"done": data.Done,
				"text": data.Text,
			},
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Modified ", result.ModifiedCount)

	return FindByID(todoID)
}

func FindByID(todoID string) (*model.Todo, error) {
	id, _ := primitive.ObjectIDFromHex(todoID)

	var todo *model.Todo
	err := getTodoRepository().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&todo)

	fmt.Println(todo)
	return todo, err
}
