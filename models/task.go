package models

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InputTask struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
}

type PatchTask struct {
	Title       *string `json:"title,omitempty" bson:"title,omitempty"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
	Status      *string `json:"status,omitempty" bson:"status,omitempty"`
}

type Task struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
	UserID      string `json:"userId" bson:"userId"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt"`
}

func AddOneTask(task *Task) error {
	collection := mongoClient.Database(DatabaseName).Collection(TaskCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, task)
	if err != nil {
		fmt.Println("Error creating task:", err)
		return err
	}

	fmt.Println("Created a task with id:", result.InsertedID)
	return nil
}

func GetTasks(userID string, page int64, perPage int64) (gin.H, error) {
	collection := mongoClient.Database(DatabaseName).Collection(TaskCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filter tasks by userId
	filter := bson.M{"userId": userID}

	// Total count
	total, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Pagination options
	skip := (page - 1) * perPage
	opts := options.Find().
		SetSkip(skip).
		SetLimit(perPage)

	// Fetch paginated data
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	// Build paginated response
	response := gin.H{
		"total":   total,
		"page":    page,
		"perPage": perPage,
		"data":    tasks,
	}

	return response, nil
}

func GetOneTask(id string) (*Task, error) {
	collection := mongoClient.Database(DatabaseName).Collection(TaskCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task Task
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		fmt.Println("Error finding task:", err)
		return nil, err
	}

	return &task, nil
}

func UpdateOneTask(id string, update bson.M) error {
	collection := mongoClient.Database(DatabaseName).Collection(TaskCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update["updatedAt"] = time.Now().Format(time.RFC3339)

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		fmt.Println("Error updating task:", err)
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no task found with id: %s", id)
	}

	return nil
}

func DeleteOneTask(id string) error {
	collection := mongoClient.Database(DatabaseName).Collection(TaskCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no task found with id: %s", id)
	}

	return nil
}
