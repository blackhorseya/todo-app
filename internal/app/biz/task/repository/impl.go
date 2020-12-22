package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/blackhorseya/todo-app/internal/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type impl struct {
	MongoClient *mongo.Client
}

// NewImpl is a constructor task of implement repository
func NewImpl(mongoClient *mongo.Client) TaskRepo {
	return &impl{MongoClient: mongoClient}
}

// CountTasks count all tasks number
func (i *impl) CountTasks() (total int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := i.MongoClient.Database("todo-db").Collection("tasks")
	res, err := coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(res), nil
}

func (i *impl) FindOne(id string) (task *entities.Task, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := i.MongoClient.Database("todo-db").Collection("tasks")
	res := coll.FindOne(ctx, bson.D{{"id", id}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	err = res.Decode(&task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// QueryTaskList handle query task list by limit and offset
func (i *impl) QueryTaskList(limit, offset int32) (tasks []*entities.Task, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := i.MongoClient.Database("todo-db").Collection("tasks")
	cur, err := coll.Find(ctx, bson.D{}, options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var task entities.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// CreateTask handle create a task
func (i *impl) CreateTask(newTask *entities.Task) (task *entities.Task, err error) {
	coll := i.MongoClient.Database("todo-db").Collection("tasks")
	_, err = coll.InsertOne(context.Background(), newTask)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

// RemoveTask handle remove a task in repository
func (i *impl) RemoveTask(id string) (count int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := i.MongoClient.Database("todo-db").Collection("tasks")
	res, err := coll.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return 0, err
	}
	if res.DeletedCount == 0 {
		return 0, fmt.Errorf("not found id: %s", id)
	}

	return int(res.DeletedCount), nil
}

func (i *impl) UpdateTask(newTask *entities.Task) (task *entities.Task, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := i.MongoClient.Database("todo-db").Collection("tasks")

	filter := bson.D{{"id", newTask.Id}}
	update := bson.D{{"$set", newTask}}
	res := coll.FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return nil, err
	}

	return newTask, nil
}
